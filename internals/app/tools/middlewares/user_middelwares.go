package middlewares

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
)

type UserMiddlewares struct {
	Http ports.HttpAdapters
}

// OnlyOwner is a middleware function that checks if the user is an owner.
// If the user is not an owner, it returns a 403 error.
func (m *UserMiddlewares) OnlyOwner(c *fiber.Ctx) error {
	role := utils.GetLocalUser(c).Role
	if role != "owner" {
		return m.Http.Forbidden(c)
	}

	return c.Next()
}