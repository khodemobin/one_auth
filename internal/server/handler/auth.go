package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/validator"
)

type AuthHandler struct {
	Logger      logger.Logger
	AuthService domain.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := new(domain.AuthRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := validator.Check(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	auth, err := h.AuthService.Login(c.Context(), req.Phone, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(auth)
}

func (h *AuthHandler) Check(c *fiber.Ctx) error {
	return c.JSON("")
}
