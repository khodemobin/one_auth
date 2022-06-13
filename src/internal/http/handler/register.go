package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/server/request"
	"github.com/khodemobin/pilo/auth/internal/service"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/helper/validator"
)

type RegisterHandler struct {
	RegisterService service.RegisterService
}

func (h *RegisterHandler) Request(c *fiber.Ctx) error {
	req := new(request.RegisterRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := validator.Check(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	err := h.RegisterService.Request(c.Context(), req.Phone, createActivity(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(helper.DefaultResponse("", "", 1))
}

func (h *RegisterHandler) Verify(c *fiber.Ctx) error {
	req := new(request.RegisterVerifyRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := validator.Check(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	auth, err := h.RegisterService.Verify(c.Context(), req.Phone, req.Code, createActivity(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createLoginCookie(c, auth)

	return c.JSON(helper.DefaultResponse(auth, "", 1))
}
