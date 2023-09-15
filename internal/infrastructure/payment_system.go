package infrastructure

import (
	"context"
	"github/cntrkilril/payment-service-emulator-golang/internal/entity"
	"github/cntrkilril/payment-service-emulator-golang/internal/service"
)

type PaymentSystemRepository struct{}

func (r *PaymentSystemRepository) FinByID(ctx context.Context, id int64) (res entity.PaymentSystem, err error) {
	// TODO connect with db or grpc-service (not included in the part of the task)
	if id != 1 {
		return entity.PaymentSystem{}, entity.ErrPaymentSystemNotFound
	}

	return entity.PaymentSystem{
		ID:   1,
		Name: "Name",
	}, nil
}

var _ service.PaymentSystemGateway = (*PaymentSystemRepository)(nil)

func NewPaymentSystemRepository() *PaymentSystemRepository {
	return &PaymentSystemRepository{}
}
