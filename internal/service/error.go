package service

import "github/cntrkilril/payment-service-emulator-golang/internal/entity"

func HandleServiceError(err error) error {
	switch err {
	case entity.ErrTransactionNotFound:
		return entity.NewError(entity.ErrTransactionNotFound.Error(), entity.ErrCodeNotFound)
	case entity.ErrPaymentSystemNotFound:
		return entity.NewError(entity.ErrPaymentSystemNotFound.Error(), entity.ErrCodeNotFound)
	case entity.ErrValidationError:
		return entity.NewError(entity.ErrValidationError.Error(), entity.ErrCodeBadRequest)
	case entity.ErrTransactionStatusCantBeCanceled:
		return entity.NewError(entity.ErrTransactionStatusCantBeCanceled.Error(), entity.ErrCodeBadRequest)
	case entity.ErrTransactionStatusCantBeUpdatedByPaymentSystem:
		return entity.NewError(entity.ErrTransactionStatusCantBeUpdatedByPaymentSystem.Error(), entity.ErrCodeBadRequest)
	default:
		return entity.NewError(entity.ErrUnknown.Error(), entity.ErrCodeInternal)
	}
}
