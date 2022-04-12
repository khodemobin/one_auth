package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h AuthHandler) UserInfo(c *fiber.Ctx) error {
	log.Println("im here")
	return nil
}
