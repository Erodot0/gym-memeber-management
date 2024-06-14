package routes

func (r *Routes) RegisterUserRoutes() {
	r.publicRoutes.Post("/users/login", r.userHandlers.Login)

	r.protectedRoutes.Post("/users", r.userHandlers.CreateUser)
	r.protectedRoutes.Post("/users/logout", r.userHandlers.Logout)
	r.protectedRoutes.Get("/users", r.userHandlers.GetUsers)
	r.protectedRoutes.Put("/users/:id", r.userHandlers.UpdateUser)
	r.protectedRoutes.Delete("/users/:id", r.userHandlers.DeleteUser)
}
