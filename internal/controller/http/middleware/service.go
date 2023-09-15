package middleware

import (
	"context"
	"github/cntrkilril/payment-service-emulator-golang/internal/entity"
)

type Service interface {
	CheckPaymentSystemIsExist(ctx context.Context, dto entity.CheckPaymentSystemDTO) error
}
