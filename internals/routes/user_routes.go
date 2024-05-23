package routes

import (
	primary "github.com/Erodot0/gym-memeber-management/internals/adapters/primary"
	secondary "github.com/Erodot0/gym-memeber-management/internals/adapters/secondary"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/services"
	"github.com/Erodot0/gym-memeber-management/internals/app/handlers"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/middlewares"
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

	userHandler := handlers.UserHandlers{
		Parser: &primary.ErrorHandler{},
		Http:   &secondary.HttpServices{},
		User:   &userServices,
	}

	userMiddleware := middlewares.UserMiddlewares{
		UserService: &userServices,
		Http:        &secondary.HttpServices{},
	}

	app.Post("/api/v1/users", userHandler.CreateUser)
	app.Post("/api/v1/users/login", userHandler.Login)
	app.Post("/api/v1/users/logout", userMiddleware.AuthorizeUser, userHandler.Logout)
}
