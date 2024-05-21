package ports

import "github.com/gofiber/fiber/v2"

type ParserAdapters interface {
	ParseData(c *fiber.Ctx, target interface{}) error
}