package routes

import (
	primary "github.com/Erodot0/gym-memeber-management/internals/adapters/primary"
	secondary "github.com/Erodot0/gym-memeber-management/internals/adapters/secondary"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/services"
	"github.com/Erodot0/gym-memeber-management/internals/app/handlers"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(app *fiber.App, db *gorm.DB, redis *redis.Client) {
	userServices := services.UserServices{
		DB: db,
		Cache: &secondary.CacheServices{
			CacheClient: redis,
		},
	}

	handler := handlers.UserHandlers{
		Parser: &primary.ErrorHandler{},
		Http:   &secondary.HttpServices{},
		User:   &userServices,
	}

	app.Post("/api/v1/users", handler.CreateUser)
	app.Post("/api/v1/users/login", handler.Login)
}
