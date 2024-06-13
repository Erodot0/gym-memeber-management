package routes

import (
	primary "github.com/Erodot0/gym-memeber-management/internals/adapters/primary"
	secondary "github.com/Erodot0/gym-memeber-management/internals/adapters/secondary"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/services"
	"github.com/Erodot0/gym-memeber-management/internals/app/handlers"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/middlewares"
)

func (r *Routes) RegisterUserRoutes() {
	userServices := services.UserServices{
		DB: r.DB,
		Cache: &secondary.CacheServices{
			CacheClient: r.Cache,
		},
	}

	userHandler := handlers.UserHandlers{
		Parser: &primary.ErrorHandler{},
		Http:   &secondary.HttpServices{},
		Services:   &userServices,
	}

	userMiddleware := middlewares.UserMiddlewares{
		UserService: &userServices,
		Http:        &secondary.HttpServices{},
	}

	r.App.Post("/api/v1/users", userMiddleware.AuthorizeUser, userMiddleware.OnlyOwner, userHandler.CreateUser)
	r.App.Post("/api/v1/users/login", userHandler.Login)
	r.App.Post("/api/v1/users/logout", userMiddleware.AuthorizeUser, userHandler.Logout)
	r.App.Get("/api/v1/users", userMiddleware.AuthorizeUser, userMiddleware.OnlyOwner, userHandler.GetUsers)
	r.App.Put("/api/v1/users/:id", userMiddleware.AuthorizeUser, userMiddleware.OnlyOwner, userHandler.UpdateUser)
	r.App.Delete("/api/v1/users/:id", userMiddleware.AuthorizeUser, userMiddleware.OnlyOwner, userHandler.DeleteUser)
}
