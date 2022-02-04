package server

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
