package service

import (
	"context"
	"github/cntrkilril/payment-service-emulator-golang/internal/controller/http/middleware"
	"github/cntrkilril/payment-service-emulator-golang/internal/entity"
)

type MiddlewareService struct {
	paymentSystemRepo PaymentSystemGateway
}

func (m *MiddlewareService) CheckPaymentSystemIsExist(ctx context.Context, dto entity.CheckPaymentSystemDTO) error {
	_, err := m.paymentSystemRepo.FinByID(ctx, dto.ID)
	if err != nil {
		return HandleServiceError(err)
	}
	return nil
}

var _ middleware.Service = (*MiddlewareService)(nil)

func NewMiddlewareService(paymentSystemRepo PaymentSystemGateway) *MiddlewareService {
	return &MiddlewareService{
		paymentSystemRepo: paymentSystemRepo,
	}
}
