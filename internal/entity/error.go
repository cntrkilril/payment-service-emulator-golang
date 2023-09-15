package entity

import "errors"

type (
	Error struct {
		msg  string
		code ErrCode
	}
	ErrCode int64
)

const (
	_ = ErrCode(iota)
	ErrCodeBadRequest
	ErrCodeInternal
	ErrCodeNotFound
	ErrCodeAccessDenied
)

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() ErrCode {
	return e.code
}

var _ error = &Error{}

func NewError(msg string, code ErrCode) *Error {
	return &Error{msg, code}
}

var (
	ErrUnknown                                       = errors.New("что-то пошло не так")
	ErrValidationError                               = errors.New("невалидные данные")
	ErrTransactionNotFound                           = errors.New("транзакция не найдена")
	ErrTransactionStatusCantBeCanceled               = errors.New("транзакцию невозможно отменить")
	ErrTransactionStatusCantBeUpdatedByPaymentSystem = errors.New("статус транзакции нельзя изменить")
	ErrPaymentSystemNotFound                         = errors.New("платежная система не найдена")
)
