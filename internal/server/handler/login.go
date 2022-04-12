package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/logger"
	"github.com/khodemobin/pilo/auth/pkg/validator"
)

type AuthHandler struct {
	Logger      logger.Logger
	AuthService domain.LoginService
}

func (h AuthHandler) Login(c *fiber.Ctx) error {
	req := new(domain.LoginRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.DefaultResponse(nil, err.Error(), 0))
	}

	errors := validator.Check(*req)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	meta := &domain.MetaData{
		Headers: c.GetRespHeaders(),
		IPs:     c.IPs(),
	}

	auth, err := h.AuthService.Login(c.Context(), req.Phone, req.Password, meta)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.DefaultResponse(nil, err.Error(), 0))
	}

	return c.JSON(helper.DefaultResponse(auth, "", 1))
}

func (h AuthHandler) Logout(c *fiber.Ctx) error {
	reqToken := c.GetRespHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	meta := &domain.MetaData{
		Headers: c.GetRespHeaders(),
		IPs:     c.IPs(),
	}

	err := h.AuthService.Logout(c.Context(), reqToken, meta)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.DefaultResponse(nil, err.Error(), 0))
	}

	return c.JSON(helper.DefaultResponse(nil, "", 1))
}
