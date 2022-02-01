package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/domain"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type Sample struct {
	Logger  logger.Logger
	Service domain.SampleService
}

func (h *Sample) Sample(c *fiber.Ctx) error {
	return c.JSON("")
}
