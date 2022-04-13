package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/server/request"
	"github.com/khodemobin/pilo/auth/internal/service"
	"github.com/khodemobin/pilo/auth/pkg/helper"
	"github.com/khodemobin/pilo/auth/pkg/validator"
)

type RegisterHandler struct {
	RegisterService service.RegisterService
}

func (h RegisterHandler) RegisterRequest(c *fiber.Ctx) error {
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

	meta := &service.MetaData{
		Headers: c.GetRespHeaders(),
		IPs:     c.IPs(),
	}

	err := h.RegisterService.RegisterRequest(c.Context(), req.Phone, meta)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(helper.DefaultResponse("", "", 1))
}

func (h RegisterHandler) RegisterVerify(c *fiber.Ctx) error {
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

	meta := &service.MetaData{
		Headers: c.GetRespHeaders(),
		IPs:     c.IPs(),
	}

	auth, err := h.RegisterService.RegisterVerify(c.Context(), req.Phone, req.Code, meta)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(helper.DefaultResponse(auth, "", 1))
}
