package ports

import "github.com/gofiber/fiber/v2"

type ParserAdapters interface {

	// ParseData parses the data from the request body into the target interface.
	//
	// Parameters:
	//   - c: the fiber.Ctx object representing the HTTP request context.
	//   - target: the interface to which the data will be parsed.
	//
	// Return:
	//   - error: if there was an error parsing the data
	ParseData(c *fiber.Ctx, target interface{}) error
}