package middleware

import (
	"net/http"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/pkg/encrypt"
	"github.com/khodemobin/pilo/auth/pkg/helper"
)

var bearerRegexp = regexp.MustCompile(`^(?:B|b)earer (\S+$)`)

func JWTChecker(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusForbidden).JSON(helper.DefaultResponse(nil, "This endpoint requires a Bearer token", 0))
	}

	matches := bearerRegexp.FindStringSubmatch(authHeader)
	if len(matches) != 2 {
		return c.Status(http.StatusForbidden).JSON(helper.DefaultResponse(nil, "This endpoint requires a Bearer token", 0))
	}

	uuid, err := encrypt.ParseJWTClaims(matches[1])
	if err != nil {
		c.ClearCookie("refresh_token")
		return c.Status(http.StatusUnauthorized).JSON(helper.DefaultResponse(nil, "", 0))
	}

	c.Locals("user_uuid", uuid)

	return c.Next()
}
