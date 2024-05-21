package configs

import (
	"log"
	"os"

	"github.com/Erodot0/gym-memeber-management/internals/routes"
	"gorm.io/gorm"
)

// Initialize sets up and starts the Fiber server.
func Initialize(db *gorm.DB) {
	app := setupFiberApp()

	newFiberLimiter(app)

	routes.RegisterUserRoutes(app)

	if err := app.Listen(":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
