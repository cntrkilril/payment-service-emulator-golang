package infrastructure

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github/cntrkilril/payment-service-emulator-golang/internal/entity"
	"github/cntrkilril/payment-service-emulator-golang/internal/service"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func (r *TransactionRepository) CountByUserID(ctx context.Context, userID int64) (res int64, err error) {
	q := `
		SELECT COUNT(*) AS count FROM transactions WHERE user_id=$1;
	`

	var count int64
	err = r.db.GetContext(ctx, &count, q, userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TransactionRepository) CountByEmail(ctx context.Context, email string) (res int64, err error) {
	q := `
		SELECT COUNT(*) AS count FROM transactions WHERE user_email=$1;
	`

	var count int64
	err = r.db.GetContext(ctx, &count, q, email)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *TransactionRepository) Save(ctx context.Context, params entity.CreateTransactionParams) (res entity.Transaction, err error) {
	q := `
		INSERT INTO transactions (user_id, user_email, amount, currency, created_at, updated_at, status)
		VALUES ($1, $2, $3, $4, now(), now(), $5)
		RETURNING id, user_id, user_email, amount, currency, created_at, updated_at, status;
		`

	err = r.db.GetContext(ctx, &res, q, params.UserID, params.UserEmail, params.Amount, params.Currency, params.Status)
	if err != nil {
		return entity.Transaction{}, err
	}

	return res, nil
}

func (r *TransactionRepository) UpdateStatus(ctx context.Context, params entity.UpdateStatusTransactionDTO) (res entity.Transaction, err error) {
	q := `
		UPDATE transactions
		SET
		    status=$2,
		    updated_at=now()
		WHERE id=$1
		RETURNING id, user_id, user_email, amount, currency, created_at, updated_at, status;
		`

	err = r.db.GetContext(ctx, &res, q, params.ID, params.Status)
	if err != nil {
		return entity.Transaction{}, err
	}

	return res, nil
}

func (r *TransactionRepository) FindByID(ctx context.Context, id int64) (res entity.Transaction, err error) {
	q := `
		SELECT id, user_id, user_email, amount, currency, created_at, updated_at, status
		FROM transactions
		WHERE id=$1;
		`

	err = r.db.GetContext(ctx, &res, q, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Transaction{}, entity.ErrTransactionNotFound
		}
		return entity.Transaction{}, err
	}

	return res, nil
}

func (r *TransactionRepository) FindByUserID(ctx context.Context, params entity.GetTransactionsByUserIDDTO) (res []entity.Transaction, err error) {
	q := `
		SELECT id, user_id, user_email, amount, currency, created_at, updated_at, status
		FROM transactions
		WHERE user_id=$1
		LIMIT $2 OFFSET $3;
		`

	err = r.db.SelectContext(ctx, &res, q, params.UserID, params.Limit, params.Offset)
	if err != nil {
		return []entity.Transaction{}, err
	}

	return res, nil
}

func (r *TransactionRepository) FindByEmail(ctx context.Context, params entity.GetTransactionsByEmailDTO) (res []entity.Transaction, err error) {
	q := `
		SELECT id, user_id, user_email, amount, currency, created_at, updated_at, status
		FROM transactions
		WHERE user_email=$1
		LIMIT $2 OFFSET $3;
		`

	err = r.db.SelectContext(ctx, &res, q, params.Email, params.Limit, params.Offset)
	if err != nil {
		return []entity.Transaction{}, err
	}

	return res, nil
}

var _ service.TransactionGateway = (*TransactionRepository)(nil)

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}
