package configs

import (
	"log"
	"os"

	"github.com/Erodot0/gym-memeber-management/internals/routes"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// Initialize sets up and starts the Fiber server.
func Initialize(db *gorm.DB, redis *redis.Client) {
	app := setupFiberApp()

	newFiberLimiter(app)

	routes := &routes.Routes{
		App:   app,
		DB:    db,
		Cache: redis,
	}

	routes.RegisterMemberRoutes()
	routes.RegisterUserRoutes()

	if err := app.Listen(":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
