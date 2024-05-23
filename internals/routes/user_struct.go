package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Routes struct {
	App   *fiber.App
	DB    *gorm.DB
	Cache *redis.Client
}
