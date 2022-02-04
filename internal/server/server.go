package server

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	handler "github.com/khodemobin/pilo/auth/internal/server/handler"

	"github.com/khodemobin/pilo/auth/internal/service"
	"github.com/khodemobin/pilo/auth/pkg/logger"
)

type Server struct {
	app         *fiber.App
	authHandler *handler.AuthHandler
}

func New(service *service.Service, isLocal bool, logger logger.Logger) *Server {
	return &Server{
		app: fiber.New(fiber.Config{
			Prefork: !isLocal,
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			},
		}),
		authHandler: &handler.AuthHandler{
			Logger:      logger,
			AuthService: service.AuthService,
		},
	}
}

func (r *Server) Start(isLocal bool, port string) error {
	if isLocal {
		r.app.Use(fiberLogger.New())
		r.app.Use(recover.New())
	}

	r.routing()
	return r.app.Listen(":" + port)
}

func (r *Server) Shutdown() error {
	return r.app.Shutdown()
}
