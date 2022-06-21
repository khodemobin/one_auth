package middleware

import (
	"fmt"
	"github.com/khodemobin/pilo/auth/internal/repository"
	"net/http"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/pkg/utils"
)

var bearerRegexp = regexp.MustCompile(`^(?:B|b)earer (\S+$)`)

func JWTChecker(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusForbidden).JSON(utils.DefaultResponse(nil, "This endpoint requires a Bearer token", 0))
	}

	matches := bearerRegexp.FindStringSubmatch(authHeader)
	if len(matches) != 2 {
		return c.Status(http.StatusForbidden).JSON(utils.DefaultResponse(nil, "This endpoint requires a Bearer token", 0))
	}

	uuid, err := utils.ParseJWTClaims(matches[1])
	if err != nil {
		c.ClearCookie("refresh_token")
		return c.Status(http.StatusUnauthorized).JSON(utils.DefaultResponse(nil, "", 0))
	}

	// exists, err := checkBlackList(matches[1])
	exists, err := repository.NewBlackListRepo().ExistsInBlackList(matches[1])
	if err != nil {
		panic(err)
	}

	if exists {
		c.ClearCookie("refresh_token")
		return c.Status(http.StatusUnauthorized).JSON(utils.DefaultResponse(nil, "", 0))
	}

	c.Locals("user_uuid", uuid)

	return c.Next()
}

func checkBlackList(token string) (bool, error) {
	value, err := app.Cache().Get(fmt.Sprintf("black_list_%s", token), nil)
	if err != nil {
		return false, err
	}

	return value != nil, nil
}
