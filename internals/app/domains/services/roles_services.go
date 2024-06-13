package services

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"gorm.io/gorm"
)

type RolesServices struct {
	DB *gorm.DB
}

func (r *RolesServices) CreateRole(role *entities.Roles) error {
	return r.DB.
		Create(role).
		Error
}

func (r *RolesServices) GetAllRoles() ([]entities.Roles, error) {
	var roles []entities.Roles
	if err := r.DB.
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RolesServices) GetRole(id uint) (*entities.Roles, error) {
	var role entities.Roles
	if err := r.DB.
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		First(&role, id).
		Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RolesServices) GetRolePermissions(roleID uint) ([]entities.Permissions, error) {
	var permissions []entities.Permissions
	if err := r.DB.
		Where("role_id = ?", roleID).
		Find(&permissions).
		Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *RolesServices) UpdateRole(id uint, role *entities.UpdateRoles) error {
	return r.DB.Model(&entities.Roles{}).Where("id = ?", id).Updates(role).Error
}

func (r *RolesServices) DeleteRole(id uint) error {
	return r.DB.Delete(&entities.Roles{}, id).Error
}
