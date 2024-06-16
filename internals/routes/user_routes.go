package routes

func (r *Routes) RegisterUserRoutes() {
	// Auth routes
	r.authRoutes.Post("/login", r.userHandlers.Login)
	r.authRoutes.Post("/logout", r.userMiddlewares.AuthorizeUser, r.userHandlers.Logout)

	r.protectedRoutes.Post("/users", r.userHandlers.CreateUser)
	r.protectedRoutes.Get("/users", r.userHandlers.GetUsers)
	r.protectedRoutes.Put("/users/:id", r.userHandlers.UpdateUser)
	r.protectedRoutes.Delete("/users/:id", r.userHandlers.DeleteUser)
}
