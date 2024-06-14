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

func NewPermissionsHandler(permissionsService ports.PermissionsServices, http ports.HttpAdapters, parser ports.ParserAdapters) *PermissionsHandler {
	return &PermissionsHandler{
		permissionsService: permissionsService,
		http:               http,
		parser:             parser,
	}
}

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
