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
	roles  ports.RolesServices
}

// NewUserHandlers creates a new UserHandlers struct.
func NewUserHandlers(parser ports.ParserAdapters, http ports.HttpAdapters, userServices ports.UserServices, rolesServices ports.RolesServices) *UserHandlers {
	return &UserHandlers{
		parser: parser,
		http:   http,
		user:   userServices,
		roles:  rolesServices,
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

	// Check if role exist
	if _, err := h.roles.GetRole(user.RoleID); err != nil {
		return h.http.BadRequest(c, "Il ruolo selezionato non esiste")
	}

	// Check if is system role
	if h.roles.IsSystemRole(user.RoleID) {
		return h.http.BadRequest(c, "Non puoi creare un utente di sistema")
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
	credentials := new(entities.UserLogin)
	if err := h.parser.ParseData(c, credentials); err != nil {
		return h.http.BadRequest(c, "Errore nella gestione dei dati")
	}

	//Validate user
	if err := credentials.Validate(); err != nil {
		return h.http.BadRequest(c, err.Error())
	}

	//Search for user
	user, err := h.user.GetUserByEmail(credentials.Email)
	if err != nil {
		return h.http.Unauthorized(c, "Email non trovata")
	}

	//Compare Password
	if err := h.user.ComparePassword(user.ID, credentials.Password); err != nil {
		return h.http.Unauthorized(c, "Password errata")
	}

	//Create Session
	if err := h.user.SetSession(c, user); err != nil {
		return h.http.InternalServerError(c, "Error creating session")
	}

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
	// Get Local User and permissions
	loggedUser := utils.GetLocalUser(c)
	permission := utils.GetLocalPermission(c)

	// Check if user exists
	user := new(entities.User)
	user.ID = utils.GetUintParam(c, "id")
	if err := u.user.GetUserById(user); err != nil {
		return u.http.NotFound(c, "Utente non trovato")
	}

	// Check if user can update
	if permission == 2 && user.ID != loggedUser.ID {
		return u.http.Forbidden(c)
	}

	// Parse data from request
	newUser := new(entities.UpdateUser)
	if err := u.parser.ParseData(c, newUser); err != nil {
		return u.http.BadRequest(c, err.Error())
	}

	// Check if role exist
	if newUser.RoleID != 0 {
		if _, err := u.roles.GetRole(newUser.RoleID); err != nil {
			return u.http.BadRequest(c, "Il ruolo selezionato non esiste")
		}
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
	// Get Local User and permissions
	loggedUser := utils.GetLocalUser(c)
	permission := utils.GetLocalPermission(c)

	user := new(entities.User)
	user.ID = utils.GetUintParam(c, "id")

	// Check if user can update
	if permission == 2 && user.ID != loggedUser.ID {
		return u.http.Forbidden(c)
	}

	// Get user
	if err := u.user.GetUserById(user); err != nil {
		return u.http.NotFound(c, "Utente non trovato")
	}

	// Delete user
	if err := u.user.DeleteUser(user); err != nil {
		return u.http.InternalServerError(c, err.Error())
	}

	// Remove session
	if err := u.user.DeleteAllSessions(c, user.ID); err != nil {
		return u.http.InternalServerError(c, err.Error())
	}

	return u.http.Success(c, nil, "Utente eliminato!")
}
