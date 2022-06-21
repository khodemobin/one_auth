package auth

import (
	"github.com/gofiber/fiber/v2"
)

type Handlers interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}
