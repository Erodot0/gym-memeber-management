package routes

import (
	primary "github.com/Erodot0/gym-memeber-management/internals/adapters/primary"
	secondary "github.com/Erodot0/gym-memeber-management/internals/adapters/secondary"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/services"
	"github.com/Erodot0/gym-memeber-management/internals/app/handlers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(app *fiber.App, db *gorm.DB) {
	userServices := services.UserServices{
		DB: db,
	}

	handler := handlers.UserHandlers{
		Parser: &primary.ErrorHandler{},
		Http:   &secondary.HttpServices{},
		User:   &userServices,
	}

	app.Post("/api/v1/users", handler.CreateUser)
	app.Post("/api/v1/users/login", handler.Login)
}
