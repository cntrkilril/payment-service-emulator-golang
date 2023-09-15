package controller

import (
	"context"
	"github/cntrkilril/payment-service-emulator-golang/internal/entity"
)

type TransactionService interface {
	Create(ctx context.Context, dto entity.CreateTransactionDTO) (entity.Transaction, error)
	UpdateStatusByPaymentSystem(ctx context.Context, dto entity.UpdateStatusTransactionDTO) (entity.Transaction, error)
	GetStatusByID(ctx context.Context, dto entity.GetTransactionByIDDTO) (entity.StatusTransaction, error)
	GetByUserID(ctx context.Context, dto entity.GetTransactionsByUserIDDTO) (entity.TransactionArray, error)
	GetByUserEmail(ctx context.Context, dto entity.GetTransactionsByEmailDTO) (entity.TransactionArray, error)
	CancelByID(ctx context.Context, dto entity.CancelTransactionByIDDTO) (entity.Transaction, error)
}
