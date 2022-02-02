package server

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
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
		// r.app.Use(recover.New())
	}

	r.routing()
	return r.app.Listen(":" + port)
}

func (r *Server) Shutdown() error {
	return r.app.Shutdown()
}

func (r *Server) routing() {
	api := r.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", r.authHandler.Login)
	v1.Post("/check", r.authHandler.Check)

	v1.Post("/register", r.authHandler.Login)
	v1.Post("/register/verify", r.authHandler.Login)
	v1.Post("/register/info", r.authHandler.Login)

	v1.Post("/recovery", r.authHandler.Login)
	v1.Post("/recovery/verify", r.authHandler.Login)
}
