package handler

import (
	"github.com/khodemobin/pilo/auth/pkg/helper/validator"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/http/request"
	"github.com/khodemobin/pilo/auth/internal/service"
	"github.com/khodemobin/pilo/auth/pkg/helper"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := new(request.LoginRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.DefaultResponse(nil, err.Error(), 0))
	}

	errors := validator.Check(*req)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	auth, err := h.AuthService.Login(c.Context(), req.Phone, req.Password, createActivity(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.DefaultResponse(nil, err.Error(), 0))
	}

	createLoginCookie(c, auth)

	return c.JSON(helper.DefaultResponse(auth, "", 1))
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	accessToken := splitToken[1]
	token := c.Cookies("refresh_token")

	err := h.AuthService.Logout(c.Context(), accessToken, token, createActivity(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.DefaultResponse(nil, err.Error(), 0))
	}

	c.ClearCookie("refresh_token")

	return c.JSON(helper.DefaultResponse(nil, "", 1))
}
