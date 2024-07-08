package handlers

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
)

type RolesHandlers struct {
	parser        ports.ParserAdapters
	http          ports.HttpAdapters
	rolesServices ports.RolesServices
}

func NewRolesHandlers(parser ports.ParserAdapters, http ports.HttpAdapters, rolesService ports.RolesServices) *RolesHandlers {
	return &RolesHandlers{
		parser:        parser,
		http:          http,
		rolesServices: rolesService,
	}
}

// CreateRole handles the creation of a new role.
func (h *RolesHandlers) CreateRole(c *fiber.Ctx) error {
	role := new(entities.Roles)
	if err := h.parser.ParseData(c, role); err != nil {
		return h.http.BadRequest(c, "Errore nella gestione dei dati")
	}

	// Validate role
	if err := role.Validate(); err != nil {
		return h.http.BadRequest(c, err.Error())
	}

	// Create role
	if err := h.rolesServices.CreateRole(role); err != nil {
		return h.http.InternalServerError(c, "Errore nel creare il ruolo")
	}

	return h.http.Success(c, []interface{}{role}, "Ruolo creato!", nil)
}

// GetAllRoles handles the retrieval of all roles.
func (h *RolesHandlers) GetAllRoles(c *fiber.Ctx) error {
	roles, err := h.rolesServices.GetAllRoles()
	if err != nil {
		return h.http.InternalServerError(c, "Errore nel recuperare i ruoli")
	}

	// Check if roles is empty
	if len(roles) == 0 {
		return h.http.NotFound(c, "Ruoli non trovati")
	}

	return h.http.Success(c, roles, "Ruoli recuperati", nil)
}

// GetRole handles the retrieval of a role by its ID.
func (h *RolesHandlers) GetRole(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "id")

	if id == 0 {
		return h.http.BadRequest(c, "Specificare l'id del ruolo")
	}

	// Get role
	role, err := h.rolesServices.GetRole(id)
	if err != nil {
		return h.http.NotFound(c, "Ruolo non trovato")
	}

	return h.http.Success(c, []interface{}{role}, "Ruolo recuperato", nil)
}

// GetRolePermissions handles the retrieval of the permissions of a role by its ID.
func (h *RolesHandlers) GetRolePermissions(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "id")

	if id == 0 {
		return h.http.BadRequest(c, "Specificare l'id del ruolo")
	}

	// Get role
	role, err := h.rolesServices.GetRolePermissions(id)
	if err != nil {
		return h.http.NotFound(c, "Ruolo non trovato")
	}

	return h.http.Success(c, role, "Permessi recuperati", nil)
}

// UpdateRole handles the update of a role.
func (h *RolesHandlers) UpdateRole(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "id")
	role := new(entities.UpdateRoles)
	if err := h.parser.ParseData(c, role); err != nil {
		return h.http.BadRequest(c, "Errore nella gestione dei dati")
	}

	// Validate role
	if err := role.Validate(); err != nil {
		return h.http.BadRequest(c, err.Error())
	}

	// Get role
	_, err := h.rolesServices.GetRole(id)
	if err != nil {
		return h.http.NotFound(c, "Ruolo non trovato")
	}

	//Update role
	if err := h.rolesServices.UpdateRole(id, role); err != nil {
		return h.http.NotFound(c, "Ruolo non trovato")
	}

	return h.http.Success(c, []interface{}{role}, "Ruolo aggiornato", nil)
}

// DeleteRole handles the deletion of a role.
func (h *RolesHandlers) DeleteRole(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "id")

	if id == 0 {
		return h.http.BadRequest(c, "Specificare l'id del ruolo")
	}

	// Get role
	_, err := h.rolesServices.GetRole(id)
	if err != nil {
		return h.http.NotFound(c, "Ruolo non trovato")
	}

	// Delete role
	if err := h.rolesServices.DeleteRole(id); err != nil {
		return h.http.NotFound(c, "Ruolo non trovato")
	}

	return h.http.Success(c, nil, "Ruolo eliminato", nil)
}
