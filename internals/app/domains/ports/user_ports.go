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

	// ComparePassword compares the input password string with the hashed password in the database.
	//
	// Parameters:
	//   - userID: the ID of the user whose password is being compared.
	//   - password: the password string to be compared with the hashed password in the database.
	//
	// Return values:
	//   - error: an error if any occurs during the comparison process.
	ComparePassword(userID uint, password string) error

	// CreateUser creates a new user in the database.
	//
	// Parameters:
	//   - user: the user entity to be created.
	//
	// Return type: error.
	CreateUser(user *entities.User) error

	// DeleteUser deletes a user from the database.
	//
	// Parameters:
	//   - u: the user entity to be deleted.
	//
	// Return type: error.
	DeleteUser(u *entities.User) error

	// GetAllUsers retrieves all users from the database.
	//
	// Return type:
	//   - []entities.User
	//
	GetAllUsers() ([]entities.User, error)

	// GetUserById retrieves a user from the database by their ID.
	//
	// Parameters:
	//   - u: a pointer to a User entity, which should have the ID field set to the desired user's ID.
	//
	// Return type: error. If the user is found, the User entity will be populated with the user's data.
	//               If the user is not found, an error will be returned.
	GetUserById(u *entities.User) error

	// GetUserForLogin retrieves a user from the database for login purposes based on the provided ID.
	//
	// Parameters:
	//   - id: the ID of the user to retrieve.
	//
	// Return type:
	//   - *entities.User
	//
	GetUserForLogin(id uint) (*entities.User, error)

	// GetUserByEmail retrieves a user from the database by their email.
	//
	// Parameters:
	//   - email: the email of the user to retrieve.
	//
	// Return type: *entities.User, error.
	GetUserByEmail(email string) (*entities.User, error)

	// UpdateUser updates a user in the database.
	//
	// Parameters:
	//   - id: the ID of the user to be updated.
	//   - u: the updated user data.
	//
	// Return type: *entities.User, error.
	UpdateUser(id uint, u *entities.UpdateUser) (*entities.User, error)

	// SetSession sets a session for a user in the database.
	//
	// Parameters:
	//   - c: the fiber.Ctx object representing the HTTP request context.
	//   - user: the user entity for which the session is being set.
	//
	// Return type: error.
	SetSession(c *fiber.Ctx, user *entities.User) error

	// GetSessionByToken retrieves a session from the database by its token.
	//
	// Parameters:
	//   - token: the token of the session to retrieve.
	//
	// Returns:
	//   - *entities.Session: the session with the given token, or nil if not found.
	//   - error: an error if the retrieval process encounters any issues.
	GetSessionByToken(token string) (*entities.Session, error)

	// DeleteSession deletes a user session from the database by its ID.
	//
	// Parameters:
	//   - c: the fiber.Ctx object representing the HTTP request context.
	//   - id: the ID of the session to be deleted.
	//
	// Return type: error.
	DeleteSession(c *fiber.Ctx, id uint) error

	// DeleteAllSessions deletes all user sessions from the database by its ID.
	//
	// Parameters:
	//   - c: the fiber.Ctx object representing the HTTP request context.
	//   - id: the ID of the user whose sessions are to be deleted.
	//
	// Return type: error.
	DeleteAllSessions(c *fiber.Ctx, id uint) error

	// CrateSystemUser creates a new system user in the database.
	CreateSystemUser() error

	// IsSystemUser checks if a user is a system user.
	IsSystemUser(id uint) bool
}
