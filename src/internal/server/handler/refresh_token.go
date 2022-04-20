package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/service"
	"github.com/khodemobin/pilo/auth/pkg/helper"
)

type RefreshTokenHandler struct {
	RefreshTokenService service.RefreshTokenService
}

func (h *RefreshTokenHandler) Refresh(c *fiber.Ctx) error {
	token := c.Cookies("refresh_token")
	log.Println(token)
	auth, err := h.RefreshTokenService.Refresh(c.Context(), token, createActivity(c))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(helper.DefaultResponse(nil, "", 0))
	}

	createLoginCookie(c, auth)

	return c.JSON(helper.DefaultResponse(auth, "", 1))
}
