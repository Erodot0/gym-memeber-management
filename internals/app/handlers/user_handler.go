package handlers

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	Parser ports.ParserAdapters
	Http ports.HttpAdapters
	User ports.UserServices
}

// CreateUser handles the creation of a new user.
//
// It takes a fiber.Ctx parameter `c` representing the HTTP request context.
// It returns an error indicating whether the user creation was successful or not.
func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := h.Parser.ParseData(c, user); err != nil {
		return h.Http.BadRequest(c, "Error parsing data")
	}

	//Validate user
	if err := user.Validate("register"); err != nil {
		return h.Http.BadRequest(c, err.Error())
	}

	// Hash password
	hashedPassword, err := h.User.EcnrypPassword(user.Password)
	if err != nil {
		return h.Http.InternalServerError(c, "Error hashing password")
	}
	user.Password = hashedPassword

	// Create user
	if err := h.User.CreateUser(user); err != nil {
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
		return h.Http.BadRequest(c, "Error parsing data")
	}

	//Validate user
	if err := user.Validate("login"); err != nil {
		return h.Http.BadRequest(c, err.Error())
	}

	provided_password := user.Password
	//Search for user
	if err := h.User.GetUserByEmail(user); err != nil {
		return h.Http.Unauthorized(c)
	}

	//Compare Password
	if err := h.User.ComparePassword(user.Password, provided_password); err != nil {
		return h.Http.Unauthorized(c)
	}

	//TODO:Create Session
	//TODO:Set Cookie

	user.RemovePassword()
	return h.Http.Success(c, []interface{}{user}, "Login successful")
}