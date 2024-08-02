package configs

import (
	"log"
	"os"

	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func setupFiberApp() *fiber.App {
	log.Println("Setting up Fiber app")
	return fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
}

func newFiberCors(app *fiber.App) {
	log.Println("Setting up CORS...")

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("ALLOW_ORIGINS"),
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
		AllowCredentials: true,
	}))
}

func newFiberLimiter(app *fiber.App) {
	log.Println("Setting up Fiber limiter...")
	app.Use("/api/v1/*", limiter.New(limiter.Config{
		Max:        100,
		Expiration: 60 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests",
			})
		},
	}))
}
