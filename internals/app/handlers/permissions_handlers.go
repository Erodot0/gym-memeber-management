package handlers

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
)

type PermissionsHandler struct {
	permissionsService ports.PermissionsServices
	http               ports.HttpAdapters
	parser             ports.ParserAdapters
}

func NewPermissionsHandler(parser ports.ParserAdapters, http ports.HttpAdapters, permissionsService ports.PermissionsServices) *PermissionsHandler {
	return &PermissionsHandler{
		http:               http,
		parser:             parser,
		permissionsService: permissionsService,
	}
}

// CreatePermission handles the creation of a new permission.
func (p *PermissionsHandler) CreatePermission(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "id")
	perm := &entities.Permissions{}
	if err := p.parser.ParseData(c, perm); err != nil {
		return err
	}

	if id != 0 {
		perm.RoleId = id
	}

	if err := p.permissionsService.ValidateNewPermission(perm); err != nil {
		return p.http.BadRequest(c, err.Error())
	}

	if err := p.permissionsService.CreatePermission(perm); err != nil {
		return p.http.InternalServerError(c, err.Error())
	}

	return p.http.Success(c, []interface{}{perm}, "Permesso creato")
}

// UpdatePermission handles the update of an existing permission.
func (p *PermissionsHandler) UpdatePermission(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "perm_id")
	if id == 0 {
		return p.http.BadRequest(c, "Specificare l'id del permesso")
	}

	perm := &entities.UpdatePermissions{}
	if err := p.parser.ParseData(c, perm); err != nil {
		return err
	}

	// Check if the permission exists
	_, err := p.permissionsService.GetPermission(id)
	if err != nil {
		return p.http.NotFound(c, "Permesso non trovato")
	}

	if err := p.permissionsService.ValidateUpdatePermission(perm); err != nil {
		return p.http.BadRequest(c, err.Error())
	}

	permission, err := p.permissionsService.UpdatePermission(id, perm)
	if err != nil {
		return p.http.InternalServerError(c, err.Error())
	}

	return p.http.Success(c, []interface{}{permission}, "Permesso aggiornato")
}

// GetPermission handles the retrieval of a permission by its ID.
func (p *PermissionsHandler) GetPermission(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "perm_id")
	if id == 0 {
		return p.http.BadRequest(c, "Specificare l'id del permesso")
	}

	permission, err := p.permissionsService.GetPermission(id)
	if err != nil {
		return p.http.NotFound(c, "Permesso non trovato")
	}

	return p.http.Success(c, []interface{}{permission}, "Permesso recuperato")
}

// GetPermissions handles the retrieval of all permissions.
func (p *PermissionsHandler) GetPermissions(c *fiber.Ctx) error {
	permissions, err := p.permissionsService.GetAllPermissions()
	if err != nil {
		return p.http.NotFound(c, "Permesso non trovato")
	}

	if len(permissions) == 0 {
		return p.http.NotFound(c, "Permessi non trovato")
	}

	return p.http.Success(c, permissions, "Permesso recuperato")
}

// DeletePermission handles the deletion of a permission by its ID.
func (p *PermissionsHandler) DeletePermission(c *fiber.Ctx) error {
	id := utils.GetUintParam(c, "perm_id")
	if id == 0 {
		return p.http.BadRequest(c, "Specificare l'id del permesso")
	}

	// Check if the permission exists
	_, err := p.permissionsService.GetPermission(id)
	if err != nil {
		return p.http.NotFound(c, "Permesso non trovato")
	}

	// Delete the permission
	if err := p.permissionsService.DeletePermission(id); err != nil {
		return p.http.NotFound(c, "Permesso non trovato")
	}

	return p.http.Success(c, nil, "Permesso eliminato")
}
