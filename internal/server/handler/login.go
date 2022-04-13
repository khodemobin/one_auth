package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/server/request"
	"github.com/khodemobin/pilo/auth/internal/service"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/validator"
)

type AuthHandler struct {
	AuthService service.LoginService
}

func (h AuthHandler) Login(c *fiber.Ctx) error {
	req := new(request.LoginRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.DefaultResponse(nil, err.Error(), 0))
	}

	errors := validator.Check(*req)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	meta := &service.MetaData{
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

	meta := &service.MetaData{
		Headers: c.GetRespHeaders(),
		IPs:     c.IPs(),
	}

	err := h.AuthService.Logout(c.Context(), reqToken, meta)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.DefaultResponse(nil, err.Error(), 0))
	}

	return c.JSON(helper.DefaultResponse(nil, "", 1))
}
