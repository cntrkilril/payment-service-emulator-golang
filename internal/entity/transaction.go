package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/shopspring/decimal"
	"math"
	"time"
)

type Transaction struct {
	ID        int64           `db:"id" json:"id"`
	UserID    int64           `db:"user_id" json:"userID"`
	UserEmail string          `db:"user_email" json:"userEmail"`
	Amount    decimal.Decimal `db:"amount" json:"amount"`
	Currency  string          `db:"currency" json:"currency"`
	CreatedAt time.Time       `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time       `db:"updated_at" json:"updatedAt"`
	Status    string          `db:"status" json:"status"`
}

type TransactionArray struct {
	Transactions []Transaction `json:"transactions"`
	Count        int64         `json:"count"`
}

type CreateTransactionDTO struct {
	UserID    int64           `json:"userID"`
	UserEmail string          `json:"userEmail"`
	Amount    decimal.Decimal `json:"amount"`
	Currency  string          `json:"currency"`
}

func (s CreateTransactionDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.UserID, validation.Required.Error("идентификатор пользователя обязателен для заполнения"),
			validation.Min(1).Error("идентификатор пользователя должен быть положительным целым числом")),
		validation.Field(&s.UserEmail, validation.Required.Error("email пользователя обязателен для заполнения"),
			is.Email.Error("неправильный формат email пользователя")),
		validation.Field(&s.Amount, validation.Required.Error("сумма транзакции обязательно для запалнения")),
		validation.Field(&s.Currency, validation.Required.Error("валюта транзакции обязательна для заполнения"),
			validation.Length(0, int(math.Inf(1))).Error("длина названия валюты должна быть больше 0")),
	)
}

type CreateTransactionParams struct {
	UserID    int64
	UserEmail string
	Amount    decimal.Decimal
	Currency  string
	Status    string
}

type UpdateStatusTransactionDTO struct {
	ID     int64  `params:"id"`
	Status string `json:"status"`
}

func (s UpdateStatusTransactionDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.ID, validation.Required.Error("идентификатор обязателен для заполнения"),
			validation.Min(1).Error("идентификатор должен быть положительным целым числом")),
		validation.Field(&s.Status, validation.Required.Error("статус обязателен для заполения"),
			validation.In(SuccessStatus.String(), FailureStatus.String()).Error("недопустимый статус")),
	)
}

type GetTransactionByIDDTO struct {
	ID int64 `params:"id"`
}

func (s GetTransactionByIDDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.ID, validation.Required.Error("идентификатор обязателен для заполнения"),
			validation.Min(1).Error("идентификатор должен быть положительным целым числом")),
	)
}

type StatusTransaction struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type CancelTransactionByIDDTO struct {
	ID int64 `params:"id"`
}

func (s CancelTransactionByIDDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.ID, validation.Required.Error("идентификатор обязателен для заполнения"),
			validation.Min(1).Error("идентификатор должен быть положительным целым числом")),
	)
}

type GetTransactionsByUserIDDTO struct {
	UserID int64 `query:"userID"`
	Limit  int64 `query:"limit"`
	Offset int64 `query:"offset"`
}

func (s GetTransactionsByUserIDDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.UserID, validation.Required.Error("идентификатор пользователя обязателен для заполнения"),
			validation.Min(1).Error("идентификатор пользователя должен быть положительным целым числом")),
		validation.Field(&s.Limit, validation.Required.Error("limit обязателен для заполнения"),
			validation.Min(1).Error("limit должен быть положительным целым числом"),
			validation.Max(100).Error("limit должен быть мешьне либо равен 100")),
		validation.Field(&s.Offset, validation.Min(0).Error("offset должен быть больше либо равен 0")),
	)
}

type GetTransactionsByEmailDTO struct {
	Email  string `query:"email"`
	Limit  int64  `query:"limit"`
	Offset int64  `query:"offset"`
}

func (s GetTransactionsByEmailDTO) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Email, validation.Required.Error("email пользователя обязателен для заполнения"),
			is.Email.Error("неправильный формат email пользователя")),
		validation.Field(&s.Limit, validation.Required.Error("limit обязателен для заполнения"),
			validation.Min(1).Error("limit должен быть положительным целым числом"),
			validation.Max(100).Error("limit должен быть мешьне либо равен 100")),
		validation.Field(&s.Offset, validation.Min(0).Error("offset должен быть больше либо равен 0")),
	)
}

type Status int64

const (
	NewStatus Status = iota
	SuccessStatus
	FailureStatus
	ErrorStatus
	CancelStatus
)

func (s Status) String() string {
	switch s {
	case NewStatus:
		return "New"
	case SuccessStatus:
		return "Success"
	case FailureStatus:
		return "Failure"
	case ErrorStatus:
		return "Error"
	case CancelStatus:
		return "Cancel"
	}
	return "unknown"
}
