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

// GetLocalMember retrieves the local member from the fiber context.
//
// Parameter: c *fiber.Ctx
// Return type: *entities.Member
func GetLocalMember(c *fiber.Ctx) *entities.Member {
	return c.Locals("member").(*entities.Member)
}

// GetLocalRole retrieves the local role from the fiber context.
//
// Parameter: c *fiber.Ctx
// Return type: *entities.Role
func GetLocalRole(c *fiber.Ctx) *entities.Roles {
	return c.Locals("role").(*entities.Roles)
}

// GetLocalPermission retrieves the local permission from the fiber context.
//
// Parameter: c *fiber.Ctx
// Return type: *entities.Permission
func GetLocalPermission(c *fiber.Ctx) uint {
	return c.Locals("permission").(uint)
}

// GetStringParam returns the value of the specified parameter from the API request.
//
// Parameters:
//   - c: The fiber.Ctx object representing the current request context.
//   - key: The key of the parameter to retrieve from the request.
//
// Returns:
//   - string: The value of the specified parameter from the API request.
func GetStringParam(c *fiber.Ctx, key string) string {
	return c.Params(key)
}

// GetUintParam returns the value of the specified parameter from the API request.
//
// Parameters:
//   - c: The fiber.Ctx object representing the current request context.
//   - key: The key of the parameter to retrieve from the request.
//
// Returns:
//   - uint: The value of the specified parameter from the API request.
func GetUintParam(c *fiber.Ctx, key string) uint {
	strVal := c.Params(key)

	// convert string to uint
	uintVal, err := StringToUint(strVal)
	if err != nil {
		return 0
	}
	return uintVal
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
