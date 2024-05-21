package routes

import (
	primary "github.com/Erodot0/gym-memeber-management/internals/adapters/primary"
	secondary "github.com/Erodot0/gym-memeber-management/internals/adapters/secondary"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/services"
	"github.com/Erodot0/gym-memeber-management/internals/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	handler := handlers.UserHandlers{
		Parser: &primary.ErrorHandler{},
		Http:   &secondary.HttpServices{},
		User:   &services.UserServices{},
	}

	app.Post("/api/v1/users", handler.CreateUser)
}
