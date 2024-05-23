package middlewares

import (
	"log"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserMiddlewares struct {
	UserService ports.UserServices
	Http ports.HttpAdapters
}

func (m *UserMiddlewares) AuthorizeUser(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authorization := c.Cookies("Authorization")
		if authorization == "" {
			// Send Unauthorized response
			return m.Http.Unauthorized(c)
		}

		// Get session from Redis
		session, err := m.UserService.GetSessionByToken(authorization)
		if err != nil {
			log.Printf("@AuthorizeUser: Error getting session: %v", err)
			return m.Http.Unauthorized(c)
		}

		// Check if it is the same IP and user agent
		if session.IPAddress != c.IP() || session.UserAgent != c.Get("User-Agent") {
			log.Printf("@AuthorizeUser: Session IP/User-Agent mismatch: %v", err)
			return m.Http.Unauthorized(c)
		}

		// Create user
		user := &entities.User{}
		user.ID = session.UserID

		// Get user from database
		if err := m.UserService.GetUserById(user); err != nil {
			log.Printf("@AuthorizeUser: Error getting user: %v", err)
			return m.Http.Unauthorized(c)
		}

		// Set session
		utils.SetLocals(c, "session", session)
		utils.SetLocals(c, "user", user)
		return c.Next()
	}
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