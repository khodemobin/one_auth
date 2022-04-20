package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/internal/model"
	"github.com/khodemobin/pilo/auth/internal/service"
	"github.com/khodemobin/pilo/auth/pkg/helper"
)

type UserHandler struct {
	UserService service.UserService
}

func (u UserHandler) UserInfo(c *fiber.Ctx) error {
	uuid := c.Locals("user_uuid")
	user, err := u.UserService.GetUser(c.Context(), uuid.(string), createActivity(c))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(helper.DefaultResponse(nil, "", 0))
	}

	data := &model.UserResource{
		Phone: user.Phone,
		UUID:  user.UUID,
	}

	return c.JSON(helper.DefaultResponse(data, "", 1))
}
