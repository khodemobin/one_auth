package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/khodemobin/pilo/auth/app"
	"github.com/khodemobin/pilo/auth/internal/http/handler"
	"github.com/khodemobin/pilo/auth/internal/service"
)

type Server struct {
	app             *fiber.App
	authHandler     handler.AuthHandler
	registerHandler handler.RegisterHandler
	userHandler     handler.UserHandler
	refreshHandler  handler.RefreshTokenHandler
}

func New(service *service.Service, isLocal bool) *Server {
	return &Server{
		app: fiber.New(fiber.Config{
			Prefork: !isLocal,
			ErrorHandler: func(ctx *fiber.Ctx, err error) error {
				app.Log().Error(err)
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			},
		}),
		authHandler: handler.AuthHandler{
			AuthService: service.AuthService,
		},
		registerHandler: handler.RegisterHandler{
			RegisterService: service.RegisterService,
		},
		userHandler: handler.UserHandler{
			UserService: service.UserService,
		},
		refreshHandler: handler.RefreshTokenHandler{
			RefreshTokenService: service.RefreshService,
		},
	}
}

func (r *Server) Start(isLocal bool, port string) error {
	if isLocal {
		r.app.Use(fiberLogger.New())
	} else {
		r.app.Use(recover.New(), compress.New())
	}

	r.app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	r.routing()
	return r.app.Listen(":" + port)
}

func (r *Server) Shutdown() error {
	return r.app.Shutdown()
}
