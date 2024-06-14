package handlers

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
)

type RolesHandlers struct {
	Parser       ports.ParserAdapters
	Http         ports.HttpAdapters
	RolesService ports.RolesServices
}

func (h *RolesHandlers) CreateRole(c *fiber.Ctx) error {
	role := new(entities.Roles)
	if err := h.Parser.ParseData(c, role); err != nil {
		return h.Http.BadRequest(c, "Errore nella gestione dei dati")
	}

	// Validate role
	if err := role.Validate(); err != nil {
		return h.Http.BadRequest(c, err.Error())
	}

	// Create role
	if err := h.RolesService.CreateRole(role); err != nil {
		return h.Http.InternalServerError(c, "Errore nel creare il ruolo")
	}

	return h.Http.Success(c, []interface{}{role}, "Ruolo creato!")
}

func (h *RolesHandlers) GetAllRoles(c *fiber.Ctx) error {
	roles, err := h.RolesService.GetAllRoles()
	if err != nil {
		return h.Http.InternalServerError(c, "Errore nel recuperare i ruoli")
	}

	// Check if roles is empty
	if len(roles) == 0 {
		return h.Http.NotFound(c, "Ruoli non trovati")
	}

	return h.Http.Success(c, roles, "Ruoli recuperati")
}

func (h *RolesHandlers) GetRole(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "id")

	if id == 0 {
		return h.Http.BadRequest(c, "Specificare l'id del ruolo")
	}

	// Get role
	role, err := h.RolesService.GetRole(id)
	if err != nil {
		return h.Http.NotFound(c, "Ruolo non trovato")
	}

	return h.Http.Success(c, []interface{}{role}, "Ruolo recuperato")
}

func (h *RolesHandlers) GerRolePermissions(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "id")

	if id == 0 {
		return h.Http.BadRequest(c, "Specificare l'id del ruolo")
	}

	// Get role
	role, err := h.RolesService.GetRolePermissions(id)
	if err != nil {
		return h.Http.NotFound(c, "Ruolo non trovato")
	}

	return h.Http.Success(c, role, "Permessi recuperati")
}

func (h *RolesHandlers) UpdateRole(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "id")
	role := new(entities.UpdateRoles)
	if err := h.Parser.ParseData(c, role); err != nil {
		return h.Http.BadRequest(c, "Errore nella gestione dei dati")
	}

	// Validate role
	if err := role.Validate(); err != nil {
		return h.Http.BadRequest(c, err.Error())
	}

	//Update role
	if err := h.RolesService.UpdateRole(id, role); err != nil {
		return h.Http.NotFound(c, "Ruolo non trovato")
	}

	return h.Http.Success(c, []interface{}{role}, "Ruolo aggiornato")
}

func (h *RolesHandlers) DeleteRole(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "id")

	if id == 0 {
		return h.Http.BadRequest(c, "Specificare l'id del ruolo")
	}

	// Get role
	_, err := h.RolesService.GetRole(id)
	if err != nil {
		return h.Http.NotFound(c, "Ruolo non trovato")
	}

	// Delete role
	if err := h.RolesService.DeleteRole(id); err != nil {
		return h.Http.NotFound(c, "Ruolo non trovato")
	}

	return h.Http.Success(c, nil, "Ruolo eliminato")
}
