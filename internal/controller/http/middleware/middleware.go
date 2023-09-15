package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github/cntrkilril/payment-service-emulator-golang/internal/entity"
)

type MdwManager struct {
	middlewareService Service
}

func (m *MdwManager) PermissionValidate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paymentSystemID := c.Locals("paymentSystemID").(int)

		err := m.middlewareService.CheckPaymentSystemIsExist(c.Context(), entity.CheckPaymentSystemDTO{ID: int64(paymentSystemID)})
		if err != nil {
			if err == entity.ErrPaymentSystemNotFound {
				return c.SendStatus(fiber.StatusForbidden)
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.Next()
	}
}

func (m *MdwManager) SessionValidate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionKey := c.Get(fiber.HeaderAuthorization)
		if sessionKey == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		//TODO add check sessionKey for payment system (not included in the part of the task)

		c.Locals("paymentSystemID", 1)

		return c.Next()
	}
}

func NewMiddlewareManager(middlewareService Service) *MdwManager {
	return &MdwManager{
		middlewareService: middlewareService,
	}
}
