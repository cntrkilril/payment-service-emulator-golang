package service

import (
	"context"
	"github/cntrkilril/payment-service-emulator-golang/internal/controller"
	"github/cntrkilril/payment-service-emulator-golang/internal/entity"
	"math/rand"
)

type TransactionService struct {
	transactionRepo TransactionGateway
}

func (s *TransactionService) Create(ctx context.Context, dto entity.CreateTransactionDTO) (entity.Transaction, error) {

	status := entity.NewStatus.String()
	if rand.Float64() <= 0.2 {
		status = entity.ErrorStatus.String()
	}

	res, err := s.transactionRepo.Save(ctx, entity.CreateTransactionParams{
		UserID:    dto.UserID,
		UserEmail: dto.UserEmail,
		Amount:    dto.Amount,
		Currency:  dto.Currency,
		Status:    status,
	})

	if err != nil {
		return entity.Transaction{}, HandleServiceError(err)
	}

	return res, nil
}

func (s *TransactionService) UpdateStatusByPaymentSystem(ctx context.Context, dto entity.UpdateStatusTransactionDTO) (entity.Transaction, error) {
	transaction, err := s.transactionRepo.FindByID(ctx, dto.ID)
	if err != nil {
		return entity.Transaction{}, HandleServiceError(err)
	}

	if transaction.Status != entity.NewStatus.String() {
		return entity.Transaction{}, HandleServiceError(entity.ErrTransactionStatusCantBeUpdatedByPaymentSystem)
	}

	res, err := s.transactionRepo.UpdateStatus(ctx, entity.UpdateStatusTransactionDTO{
		ID:     dto.ID,
		Status: dto.Status,
	})
	if err != nil {
		return entity.Transaction{}, HandleServiceError(err)
	}

	return res, nil
}

func (s *TransactionService) GetStatusByID(ctx context.Context, dto entity.GetTransactionByIDDTO) (entity.StatusTransaction, error) {
	res, err := s.transactionRepo.FindByID(ctx, dto.ID)
	if err != nil {
		return entity.StatusTransaction{}, HandleServiceError(err)
	}

	return entity.StatusTransaction{
		ID:     res.ID,
		Status: res.Status,
	}, nil
}

func (s *TransactionService) GetByUserID(ctx context.Context, dto entity.GetTransactionsByUserIDDTO) (entity.TransactionArray, error) {
	res, err := s.transactionRepo.FindByUserID(ctx, dto)
	if err != nil {
		return entity.TransactionArray{}, HandleServiceError(err)
	}

	count, err := s.transactionRepo.CountByUserID(ctx, dto.UserID)
	if err != nil {
		return entity.TransactionArray{}, HandleServiceError(err)
	}

	return entity.TransactionArray{
		Transactions: res,
		Count:        count,
	}, nil
}

func (s *TransactionService) GetByUserEmail(ctx context.Context, dto entity.GetTransactionsByEmailDTO) (entity.TransactionArray, error) {
	res, err := s.transactionRepo.FindByEmail(ctx, dto)
	if err != nil {
		return entity.TransactionArray{}, HandleServiceError(err)
	}

	count, err := s.transactionRepo.CountByEmail(ctx, dto.Email)
	if err != nil {
		return entity.TransactionArray{}, HandleServiceError(err)
	}

	return entity.TransactionArray{
		Transactions: res,
		Count:        count,
	}, nil
}

func (s *TransactionService) CancelByID(ctx context.Context, dto entity.CancelTransactionByIDDTO) (entity.Transaction, error) {
	transaction, err := s.transactionRepo.FindByID(ctx, dto.ID)
	if err != nil {
		return entity.Transaction{}, HandleServiceError(err)
	}

	if transaction.Status == entity.FailureStatus.String() || transaction.Status == entity.SuccessStatus.String() || transaction.Status == entity.CancelStatus.String() {
		return entity.Transaction{}, HandleServiceError(entity.ErrTransactionStatusCantBeCanceled)
	}

	res, err := s.transactionRepo.UpdateStatus(ctx, entity.UpdateStatusTransactionDTO{ID: dto.ID, Status: entity.CancelStatus.String()})
	if err != nil {
		return entity.Transaction{}, HandleServiceError(err)
	}

	return res, nil
}

var _ controller.TransactionService = (*TransactionService)(nil)

func NewTransactionService(transactionRepo TransactionGateway) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
	}
}
