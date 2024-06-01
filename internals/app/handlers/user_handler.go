package handlers

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	Parser   ports.ParserAdapters
	Http     ports.HttpAdapters
	Services ports.UserServices
}

// CreateUser handles the creation of a new user.
//
// It takes a fiber.Ctx parameter `c` representing the HTTP request context.
// It returns an error indicating whether the user creation was successful or not.
func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := h.Parser.ParseData(c, user); err != nil {
		return h.Http.BadRequest(c, "Errore nella gestione dei dati")
	}

	//Validate user
	if err := user.Validate("register"); err != nil {
		return h.Http.BadRequest(c, err.Error())
	}

	// Hash password
	hashedPassword, err := h.Services.EcnrypPassword(user.Password)
	if err != nil {
		return h.Http.InternalServerError(c, "Error hashing password")
	}
	user.Password = hashedPassword

	// Create user
	if err := h.Services.CreateUser(user); err != nil {
		return h.Http.InternalServerError(c, "Error creating user")
	}

	user.RemovePassword()
	return h.Http.Success(c, []interface{}{user}, "User created")
}

// Login handles the login process for a user.
//
// It takes a fiber.Ctx parameter `c` representing the HTTP request context.
// It returns an error indicating whether the login was successful or not.
func (h *UserHandlers) Login(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := h.Parser.ParseData(c, user); err != nil {
		return h.Http.BadRequest(c, "Errore nella gestione dei dati")
	}

	//Validate user
	if err := user.Validate("login"); err != nil {
		return h.Http.BadRequest(c, err.Error())
	}

	provided_password := user.Password
	//Search for user
	if err := h.Services.GetUserByEmail(user); err != nil {
		return h.Http.Unauthorized(c)
	}

	//Compare Password
	if err := h.Services.ComparePassword(user.Password, provided_password); err != nil {
		return h.Http.Unauthorized(c)
	}

	//Create Session
	if err := h.Services.SetSession(c, user); err != nil {
		return h.Http.InternalServerError(c, "Error creating session")
	}

	user.RemovePassword()
	return h.Http.Success(c, []interface{}{user}, "Login successful")
}

// Logout handles the logout process for a user.
//
// It takes a fiber.Ctx parameter `c` representing the HTTP request context.
// It returns an error indicating whether the logout was successful or not.
func (h *UserHandlers) Logout(c *fiber.Ctx) error {
	user := utils.GetLocalUser(c)

	if err := h.Services.DeleteSession(c, user.ID); err != nil {
		return h.Http.InternalServerError(c, "Error deleting session")
	}

	// Clear the cookie
	c.ClearCookie("Authorization")

	return h.Http.Success(c, nil, "Logout successful")
}

// GetUsers handles the retrieval of all users.
//
// It takes a fiber.Ctx parameter `c` representing the HTTP request context.
// It returns an error indicating whether the retrieval was successful or not.
func (u *UserHandlers) GetUsers(c *fiber.Ctx) error {
	users, err := u.Services.GetAllUsers()
	if err != nil {
		return u.Http.InternalServerError(c, err.Error())
	}
	return u.Http.Success(c, users, "Users retrieved")
}

// DeleteUser handles the deletion of a user.
//
// It takes a fiber.Ctx parameter `c` representing the HTTP request context.
// It returns an error indicating whether the user deletion was successful or not.
//
// The function retrieves the user from the Firebase local context, deletes the user,
// and removes all sessions associated with the user. If any error occurs during
// the process, an internal server error is returned. Otherwise, a success message
// is returned indicating that the user was deleted successfully.
func (u *UserHandlers) DeleteUser(c *fiber.Ctx) error {
	user := new(entities.User)
	user.ID = utils.GetUintParam(c, "id")

	// Get user
	if err := u.Services.GetUserById(user); err != nil {
		return u.Http.NotFound(c, "User not found")
	}

	// Check if user is an owner
	if user.Role == "owner" {
		return u.Http.BadRequest(c, "You can't delete the owner")
	}

	// Delete user
	if err := u.Services.DeleteUser(user); err != nil {
		return u.Http.InternalServerError(c, err.Error())
	}

	// Remove session
	if err := u.Services.DeleteAllSessions(c, user.ID); err != nil {
		return u.Http.InternalServerError(c, err.Error())
	}

	return u.Http.Success(c, nil, "User deleted successfully!")
}
