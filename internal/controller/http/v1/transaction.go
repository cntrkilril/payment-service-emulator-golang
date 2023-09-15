package v1

import (
	"github.com/gofiber/fiber/v2"
	"github/cntrkilril/payment-service-emulator-golang/internal/controller"
	"github/cntrkilril/payment-service-emulator-golang/internal/controller/http/middleware"
	"github/cntrkilril/payment-service-emulator-golang/internal/entity"
)

type TransactionHandler struct {
	transactionService controller.TransactionService
	middlewareManager  *middleware.MdwManager
}

func (h *TransactionHandler) createTransaction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.CreateTransactionDTO

		if err := c.BodyParser(&p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(entity.ErrValidationError),
			)
		}

		if err := p.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(err),
			)
		}

		if !p.Amount.IsPositive() {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(entity.NewError("сумма транзакции должна быть положительной", entity.ErrCodeBadRequest)),
			)
		}

		res, err := h.transactionService.Create(c.Context(), p)
		if err != nil {
			return HandleError(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{
				"transaction": res,
			},
			fiber.Map{
				"success": true,
			},
		))
	}
}

func (h *TransactionHandler) updateStatusTransactionByPaymentSystem() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.UpdateStatusTransactionDTO

		if err := c.BodyParser(&p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(entity.ErrValidationError),
			)
		}

		if err := c.ParamsParser(&p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(err),
			)
		}

		if err := p.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(err),
			)
		}
		res, err := h.transactionService.UpdateStatusByPaymentSystem(c.Context(), p)
		if err != nil {
			return HandleError(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{
				"transaction": res,
			},
			fiber.Map{
				"success": true,
			},
		))
	}
}

func (h *TransactionHandler) getStatusTransactionByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.GetTransactionByIDDTO

		if err := c.ParamsParser(&p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(entity.ErrValidationError),
			)
		}

		if err := p.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(err),
			)
		}
		res, err := h.transactionService.GetStatusByID(c.Context(), p)
		if err != nil {
			return HandleError(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{
				"transaction": res,
			},
			fiber.Map{
				"success": true,
			},
		))
	}
}

func (h *TransactionHandler) getTransactionsByUserID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.GetTransactionsByUserIDDTO

		if err := c.QueryParser(&p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(entity.ErrValidationError),
			)
		}

		if err := p.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(err),
			)
		}

		res, err := h.transactionService.GetByUserID(c.Context(), p)
		if err != nil {
			return HandleError(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{
				"transactions": res.Transactions,
			},
			fiber.Map{
				"success": true,
				"count":   res.Count,
			},
		))
	}
}

func (h *TransactionHandler) getTransactionsByEmail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.GetTransactionsByEmailDTO

		if err := c.QueryParser(&p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(entity.ErrValidationError),
			)
		}

		if err := p.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(err),
			)
		}
		res, err := h.transactionService.GetByUserEmail(c.Context(), p)
		if err != nil {
			return HandleError(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{
				"transactions": res.Transactions,
			},
			fiber.Map{
				"success": true,
				"count":   res.Count,
			},
		))
	}
}

func (h *TransactionHandler) cancelTransactionByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var p entity.CancelTransactionByIDDTO

		if err := c.ParamsParser(&p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(entity.ErrValidationError),
			)
		}

		if err := p.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				newErrResp(err),
			)
		}
		res, err := h.transactionService.CancelByID(c.Context(), p)
		if err != nil {
			return HandleError(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(newResp(
			fiber.Map{
				"transaction": res,
			},
			fiber.Map{
				"success": true,
			},
		))
	}
}

func (h *TransactionHandler) Register(r fiber.Router) {
	r.Post("",
		h.createTransaction())
	r.Patch("status/:id",
		h.middlewareManager.SessionValidate(),
		h.middlewareManager.PermissionValidate(),
		h.updateStatusTransactionByPaymentSystem())
	r.Get("status/:id",
		h.getStatusTransactionByID())
	r.Get("user-id",
		h.getTransactionsByUserID())
	r.Get("email",
		h.getTransactionsByEmail())
	r.Patch("cancel/:id",
		h.cancelTransactionByID())
}

func NewTransactionHandler(
	transactionService controller.TransactionService,
	middlewareManager *middleware.MdwManager,
) *TransactionHandler {
	return &TransactionHandler{
		transactionService: transactionService,
		middlewareManager:  middlewareManager,
	}
}
