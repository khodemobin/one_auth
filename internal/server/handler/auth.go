package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/internal/server/request"
	"github.com/khodemobin/pilo/auth/internal/server/response"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type AuthHandler struct {
	Logger      logger.Logger
	AuthService domain.AuthService
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := new(request.LoginRequest)
	code, err := encrypt.Hash("123456")
	fmt.Println(code)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := request.Validate(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	auth, err := h.AuthService.Login(c.Context(), req.Phone, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res := response.NewAuthResource(auth)

	return c.JSON(res)
}

func (h *AuthHandler) Check(c *fiber.Ctx) error {
	return c.JSON("")
}
