package server

import "github.com/khodemobin/pilo/auth/internal/server/middleware"

func (r *Server) routing() {
	api := r.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", r.authHandler.Login)
	v1.Post("/logout", r.authHandler.Logout)

	v1.Post("/register", r.registerHandler.RegisterRequest)
	v1.Post("/register/verify", r.registerHandler.RegisterVerify)

	v1.Post("/recovery", r.authHandler.Login)
	v1.Post("/recovery/verify", r.authHandler.Login)

	v1.Post("/refresh_token", r.refreshHandler.RefreshToken)

	auth := v1.Use(middleware.JWTChecker)
	auth.Get("/me", r.userHandler.UserInfo)
}
