package http

func (r *Server) routing() {
	//TODO add activity logger middleware
	api := r.app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", r.authHandler.Login)

	//v1.Post("/register", r.registerHandler.Request)
	//v1.Post("/register/verify", r.registerHandler.Verify)

	//v1.Post("/recovery", r.authHandler.Login)
	//v1.Post("/recovery/verify", r.authHandler.Login)

	//v1.Post("/refresh_token", r.refreshHandler.Refresh)

	//auth := v1.Use(middleware.JWTChecker)
	//auth.Get("/me", r.userHandler)
	//auth.Post("/update", r.userHandler.Update)
	//auth.Post("/logout", r.authHandler.Logout)
}
