package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type AuthHandler struct {
	Logger  logger.Logger
	Service domain.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	return c.JSON("")
}

func (h *AuthHandler) Check(c *fiber.Ctx) error {
	return c.JSON("")
}
