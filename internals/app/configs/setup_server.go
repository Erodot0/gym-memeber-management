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
	log.Println("Setting up server...")
	app := setupFiberApp()

	newFiberCors(app)
	newFiberLimiter(app)

	routes := routes.NewRoutes(app, db, redis)

	routes.RegisterMemberRoutes()
	routes.RegisterUserRoutes()
	routes.RegisterRolesRoutes()
	routes.RegisterPermissionsRoutes()

	if err := app.Listen(":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	log.Printf("Server started on port %s", os.Getenv("SERVER_PORT"))
}
