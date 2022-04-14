package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/pkg/user_agent"
)

func createActivity(c *fiber.Ctx) *model.Activity {
	var a model.Activity
	u := user_agent.Parse(string(c.Request().Header.UserAgent()))

	a.Action = string(c.Request().Header.Method())
	a.IP = c.Context().RemoteIP().String()
	a.Path = c.Request().URI().String()
	a.Operation = u.OS
	a.Version = u.OSVersion
	a.Headers = c.Request().Header.String()

	return &a
}
