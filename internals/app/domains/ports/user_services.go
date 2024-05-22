package ports

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/gofiber/fiber/v2"
)

type UserServices interface {
	// EcnrypPassword generates a hashed password from the input password string using bcrypt.
	//
	// Parameters:
	//   - password: the password string to be hashed.
	//
	// Return values:
	//   - string: the hashed password.
	//   - error: an error if any occurs during the hashing process.
	EcnrypPassword(password string) (string, error)
	// ComparePassword compares a hashed password with a plaintext password.
	//
	// Parameters:
	//   - hashedPassword: the hashed password to compare.
	//   - password: the plaintext password to compare.
	//
	// Return type: error.
	ComparePassword(hashedPassword, password string) error
	// CreateUser creates a new user in the database.
	//
	// Parameters:
	//   - user: the user entity to be created.
	//
	// Return type: error.
	CreateUser(user *entities.User) error
	// GetUserByEmail retrieves a user from the database by their email.
	//
	// Parameters:
	//   - user: the user entity to be retrieved, it must contain the email.
	//
	// Return type: error.
	GetUserByEmail(user *entities.User) error
	// SetSession sets a session for a user in the database.
	//
	// Parameters:
	//   - c: the fiber.Ctx object representing the HTTP request context.
	//   - user: the user entity for which the session is being set.
	//
	// Return type: error.
	SetSession(c *fiber.Ctx, user *entities.User) error
	GetSessionByToken(token string) (*entities.Session, error)
}