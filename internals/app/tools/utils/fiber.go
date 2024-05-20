package utils

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/gofiber/fiber/v2"
)

// SetLocals sets the local key-value pair in the fiber context.
//
// Parameters:
//
//	c *fiber.Ctx: the fiber context
//	key string: the key for the local data
//	data interface{}: the data to be stored as local
func SetLocals(c *fiber.Ctx, key string, data interface{}) {
	c.Locals(key, data)
}

// GetLocalUser returns the user from the given fiber.Ctx.
//
// It takes a parameter 'c' of type *fiber.Ctx.
// It returns a pointer to an entities.User.
func GetLocalUser(c *fiber.Ctx) *entities.User {
	return c.Locals("user").(*entities.User)
}

// GetLocalSession retrieves the local session from the fiber context.
//
// Parameter: c *fiber.Ctx
// Return type: *entities.Session
func GetLocalSession(c *fiber.Ctx) *entities.Session {
	return c.Locals("session").(*entities.Session)
}

// GetApiParam returns the value of the specified parameter from the API request.
//
// Parameters:
//   - c: The fiber.Ctx object representing the current request context.
//   - key: The key of the parameter to retrieve from the request.
//
// Returns:
//   - string: The value of the specified parameter from the API request.
func GetApiParam(c *fiber.Ctx, key string) string {
	return c.Params(key)
}

// GetBaseURL retrieves the base URL from the fiber context.
//
// Parameter:
//   - c: The fiber.Ctx object representing the current request context.
// Return type:
//   - string: The base URL from the fiber context.
func GetBaseURL(c *fiber.Ctx) string {
	return c.BaseURL()
}
