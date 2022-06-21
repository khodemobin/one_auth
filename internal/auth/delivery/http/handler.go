package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/config"
	"github.com/khodemobin/pilo/auth/internal/auth"
	"github.com/khodemobin/pilo/auth/internal/http/request"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/utils"
	"time"
)

// Auth handlers
type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
	logger logger.Logger
}

func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase, log logger.Logger) auth.Handlers {
	return &authHandlers{cfg: cfg, authUC: authUC, sessUC: sessUC, logger: log}
}

func (h authHandlers) Login(c *fiber.Ctx) error {
	var req request.LoginRequest

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.DefaultResponse(nil, err.Error(), 0))
	}

	errors := utils.ValidateCheck(&req)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	result, err := h.authUC.Login(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.DefaultResponse(nil, err.Error(), 0))
	}

	h.createLoginCookie(c, result)

	return c.JSON(utils.DefaultResponse(result, "", 1))
}

func (h authHandlers) Logout(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h authHandlers) RefreshToken(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h authHandlers) createLoginCookie(c *fiber.Ctx, auth *auth.Auth) {
	c.ClearCookie("refresh_token")
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    auth.RefreshToken.Token,
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(720 * time.Hour), // 30 day
	})
}
