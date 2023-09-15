package app

import (
	"github.com/jmoiron/sqlx"
	"github/cntrkilril/payment-service-emulator-golang/internal/controller/http/middleware"
	v1 "github/cntrkilril/payment-service-emulator-golang/internal/controller/http/v1"
	"github/cntrkilril/payment-service-emulator-golang/internal/infrastructure"
	"github/cntrkilril/payment-service-emulator-golang/internal/service"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Run() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	atom := zap.NewAtomicLevel()
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		atom,
	)
	logger := zap.New(zapCore)
	defer func(logger *zap.Logger) {
		err = logger.Sync()
		if err != nil {
			return
		}
	}(logger)

	l := logger.Sugar()
	atom.SetLevel(zapcore.Level(*cfg.Logger.Level))
	l.Infof("logger initialized successfully")

	f := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	f.Use(fiberzap.New(fiberzap.Config{
		Logger: logger,
	}))
	f.Use(cors.New(cors.Config{
		AllowHeaders: "*",
	}))

	l.Infof("fiber initialized successfully")

	// db
	db, err := sqlx.Connect("pgx", cfg.Postgres.ConnString)
	if err != nil {
		l.Error(err)
		return
	}

	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			l.Error(err)
			return
		}
	}(db)

	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Second)
	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.Postgres.ConnMaxIdleTime * time.Second)

	err = db.Ping()
	if err != nil {
		l.Error(err)
		return
	}

	l.Debug("Connected to PostgreSQL")

	// infrastructures
	transactionRepo := infrastructure.NewTransactionRepository(db)
	paymentSystemRepo := infrastructure.NewPaymentSystemRepository()

	// services
	middlewareService := service.NewMiddlewareService(paymentSystemRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	// controllers
	middlewareManager := middleware.NewMiddlewareManager(middlewareService)
	transactionHandler := v1.NewTransactionHandler(transactionService, middlewareManager)

	// groups
	apiGroup := f.Group("api")
	transactionGroup := apiGroup.Group("transaction")

	transactionHandler.Register(transactionGroup)

	go func() {
		err = f.Listen(net.JoinHostPort(cfg.HTTP.Host, cfg.HTTP.Port))
		if err != nil {
			l.Fatal(err.Error())
		}
	}()

	l.Debug("Started HTTP server")

	l.Debug("Application has started")

	exit := make(chan os.Signal, 2)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	l.Info("Application has been shut down")

}
