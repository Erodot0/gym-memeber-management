package services

import (
	"fmt"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"gorm.io/gorm"
)

type PermissionsService struct {
	db *gorm.DB
}

func NewPermissionsService(db *gorm.DB) *PermissionsService {
	return &PermissionsService{
		db: db,
	}
}

func (p *PermissionsService) ValidateNewPermission(perm *entities.Permissions) error {
	// Check if the role exists
	if perm.RoleId == 0 {
		return fmt.Errorf("il ruolo è obbligatorio")
	}

	// Check role in the database
	role := &entities.Roles{}
	if err := p.db.First(role, perm.RoleId).Error; err != nil {
		return fmt.Errorf("il ruolo non è valido")
	}

	// Check if the table exists
	if perm.Table == "" {
		return fmt.Errorf("la tabella è obbligatoria")
	}

	// Check if the action is valid
	if perm.Create > 1 || perm.Update > 2 || perm.Read > 2 || perm.Delete > 2 {
		return fmt.Errorf("assegnare i permessi correttamente")
	}

	// Check if the permission already exists
	exists, err := p.CheckPermissionExists(perm.Table, perm.RoleId)
	if err != nil {
		return fmt.Errorf("errore nel controllo dell'esistenza del permesso")
	}
	if exists {
		return fmt.Errorf("i permessi per questa tabella e ruolo sono già presenti")
	}

	return nil
}

func (p *PermissionsService) ValidateUpdatePermission(perm *entities.UpdatePermissions) error {
	// Check if the action is valid
	if perm.Create > 1 || perm.Update > 2 || perm.Read > 2 || perm.Delete > 2 {
		return fmt.Errorf("assegnare i permessi correttamente")
	}

	return nil
}

func (p *PermissionsService) CreatePermission(perm *entities.Permissions) error {
	return p.db.Create(perm).Error
}

func (p *PermissionsService) GetPermission(id uint) (*entities.Permissions, error) {
	perm := &entities.Permissions{}
	return perm, p.db.First(perm, id).Error
}

func (p *PermissionsService) GetAllPermissions() ([]entities.Permissions, error) {
	perms := []entities.Permissions{}
	return perms, p.db.Find(&perms).Error
}

func (p *PermissionsService) GetPermissionsByRole(roleId uint) ([]entities.Permissions, error) {
	perms := []entities.Permissions{}
	return perms, p.db.Where("role_id = ?", roleId).Find(&perms).Error
}

func (p *PermissionsService) GetPermissionsByTable(table string) ([]entities.Permissions, error) {
	perms := []entities.Permissions{}
	return perms, p.db.Where("table = ?", table).Find(&perms).Error
}

func (p *PermissionsService) UpdatePermission(id uint, perm *entities.UpdatePermissions) (*entities.Permissions, error) {
	if err := p.db.
		Model(&entities.Permissions{}).
		Where("id = ?", id).
		Updates(perm).Error; err != nil {
		return nil, err
	}

	return p.GetPermission(id)
}

func (p *PermissionsService) DeletePermission(id uint) error {
	return p.db.Delete(&entities.Permissions{}, id).Error
}

func (p *PermissionsService) HasPermission(table string, roleId uint, action string) (bool, error) {
	perm := &entities.Permissions{}
	return p.db.
		Where(
			"table = ? AND role_id = ? AND action = ?",
			table, roleId, action).
		First(perm).
		Error == nil, nil
}

func (p *PermissionsService) CheckPermissionExists(table string, roleId uint) (bool, error) {
	perm := &entities.Permissions{}
	return p.db.
		Where(
			"table = ? AND role_id = ?",
			table, roleId).
		First(perm).
		Error == nil, nil
}
