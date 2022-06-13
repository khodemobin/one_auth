package handler

import (
	"github.com/khodemobin/pilo/auth/pkg/helper/user_agent"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/service"
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

func createLoginCookie(c *fiber.Ctx, auth *service.Auth) {
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
