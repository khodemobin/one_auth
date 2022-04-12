package handler

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/vmihailenco/taskq/v3"
)

func (h AuthHandler) UserInfo(c *fiber.Ctx) error {
	t := taskq.RegisterTask(&taskq.TaskOptions{
		Handler: func(name string) error {
			fmt.Println("Hello", name)
			return nil
		},
	}).WithArgs(context.Background())

	// m := t.WithArgs(context.Background(), "World")

	app.Queue().Add(t, m)

	return nil
}
