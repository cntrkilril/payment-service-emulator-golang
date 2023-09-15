package service

import (
	"context"
	"github/cntrkilril/payment-service-emulator-golang/internal/entity"
)

type (
	TransactionGateway interface {
		Save(ctx context.Context, params entity.CreateTransactionParams) (res entity.Transaction, err error)
		UpdateStatus(ctx context.Context, params entity.UpdateStatusTransactionDTO) (res entity.Transaction, err error)
		FindByID(ctx context.Context, id int64) (res entity.Transaction, err error)
		FindByUserID(ctx context.Context, params entity.GetTransactionsByUserIDDTO) (res []entity.Transaction, err error)
		CountByUserID(ctx context.Context, userID int64) (res int64, err error)
		FindByEmail(ctx context.Context, params entity.GetTransactionsByEmailDTO) (res []entity.Transaction, err error)
		CountByEmail(ctx context.Context, email string) (res int64, err error)
	}
	PaymentSystemGateway interface {
		FinByID(ctx context.Context, id int64) (res entity.PaymentSystem, err error)
	}
)
