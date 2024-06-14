package handlers

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	parser ports.ParserAdapters
	http   ports.HttpAdapters
	user   ports.UserServices
}

// NewUserHandlers creates a new UserHandlers struct.
func NewUserHandlers(parser ports.ParserAdapters, http ports.HttpAdapters, services ports.UserServices) *UserHandlers {
	return &UserHandlers{
		parser: parser,
		http:   http,
		user:   services,
	}
}

// CreateUser handles the creation of a new user.
func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := h.parser.ParseData(c, user); err != nil {
		return h.http.BadRequest(c, "Errore nella gestione dei dati")
	}

	//Validate user
	if err := user.Validate(); err != nil {
		return h.http.BadRequest(c, err.Error())
	}

	// Hash password
	hashedPassword, err := h.user.EcnrypPassword(user.Password)
	if err != nil {
		return h.http.InternalServerError(c, "Error hashing password")
	}
	user.Password = hashedPassword

	// Create user
	if err := h.user.CreateUser(user); err != nil {
		return h.http.InternalServerError(c, "Error creating user")
	}

	user.RemovePassword()
	return h.http.Success(c, []interface{}{user}, "User created")
}

// Login handles the login process for a user.
func (h *UserHandlers) Login(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := h.parser.ParseData(c, user); err != nil {
		return h.http.BadRequest(c, "Errore nella gestione dei dati")
	}

	//Validate user
	if err := user.ValidateLogin(); err != nil {
		return h.http.BadRequest(c, err.Error())
	}

	provided_password := user.Password
	//Search for user
	if err := h.user.GetUserByEmail(user); err != nil {
		return h.http.Unauthorized(c)
	}

	//Compare Password
	if err := h.user.ComparePassword(user.Password, provided_password); err != nil {
		return h.http.Unauthorized(c)
	}

	//Create Session
	if err := h.user.SetSession(c, user); err != nil {
		return h.http.InternalServerError(c, "Error creating session")
	}

	user.RemovePassword()
	return h.http.Success(c, []interface{}{user}, "Login successful")
}

// Logout handles the logout process for a user.
func (h *UserHandlers) Logout(c *fiber.Ctx) error {
	user := utils.GetLocalUser(c)

	if err := h.user.DeleteSession(c, user.ID); err != nil {
		return h.http.InternalServerError(c, "Error deleting session")
	}

	// Clear the cookie
	c.ClearCookie("Authorization")

	return h.http.Success(c, nil, "Logout successful")
}

// GetUsers handles the retrieval of all users.
func (u *UserHandlers) GetUsers(c *fiber.Ctx) error {
	users, err := u.user.GetAllUsers()
	if err != nil {
		return u.http.InternalServerError(c, err.Error())
	}
	return u.http.Success(c, users, "Utenti recuperati correttamente")
}

// UpdateUser handles the update of a user.
func (u *UserHandlers) UpdateUser(c *fiber.Ctx) error {
	// Check if user exists
	user := new(entities.User)
	user.ID = utils.GetUintParam(c, "id")
	if err := u.user.GetUserById(user); err != nil {
		return u.http.NotFound(c, "Utente non trovato")
	}

	newUser := new(entities.UpdateUser)
	if err := u.parser.ParseData(c, newUser); err != nil {
		return u.http.BadRequest(c, err.Error())
	}

	// update user
	user, err := u.user.UpdateUser(user.ID, newUser)
	if err != nil {
		return u.http.InternalServerError(c, err.Error())
	}

	return u.http.Success(c, []interface{}{user}, "User updated")
}

// DeleteUser handles the deletion of a user.
func (u *UserHandlers) DeleteUser(c *fiber.Ctx) error {
	user := new(entities.User)
	user.ID = utils.GetUintParam(c, "id")

	// Get user
	if err := u.user.GetUserById(user); err != nil {
		return u.http.NotFound(c, "User not found")
	}

	// Check if user is an owner
	if user.Role.Name == "owner" {
		return u.http.BadRequest(c, "You can't delete the owner")
	}

	// Delete user
	if err := u.user.DeleteUser(user); err != nil {
		return u.http.InternalServerError(c, err.Error())
	}

	// Remove session
	if err := u.user.DeleteAllSessions(c, user.ID); err != nil {
		return u.http.InternalServerError(c, err.Error())
	}

	return u.http.Success(c, nil, "User deleted successfully!")
}
