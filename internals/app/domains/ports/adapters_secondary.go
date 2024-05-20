package ports

import "github.com/gofiber/fiber/v2"

// HttpResponses defines methods for handling HTTP responses
type HttpAdapters interface {
	Success(c *fiber.Ctx, data interface{}, message string) error
	BadRequest(c *fiber.Ctx, message string) error
	Unauthorized(c *fiber.Ctx) error
	Forbidden(c *fiber.Ctx) error
	NotFound(c *fiber.Ctx, message string) error
	InternalServerError(c *fiber.Ctx, message string) error
	WithFile(c *fiber.Ctx, pathToFile string) error
}