package services

import (
	"errors"
	"log"
	"os"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"gorm.io/gorm"
)

type RolesServices struct {
	db *gorm.DB
}

func NewRolesServices(db *gorm.DB) *RolesServices {
	return &RolesServices{
		db: db,
	}
}

func (r *RolesServices) CreateRole(role *entities.Roles) error {
	return r.db.
		Create(role).
		Error
}

func (r *RolesServices) GetAllRoles() ([]entities.Roles, error) {
	systemRoleName := os.Getenv("SYS_ROLE_NAME")

	var roles []entities.Roles
	if err := r.db.
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Where("name != ?", systemRoleName).
		Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RolesServices) GetRole(id uint) (*entities.Roles, error) {
	systemRoleName := os.Getenv("SYS_ROLE_NAME")

	var role entities.Roles
	if err := r.db.
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Where("name != ?", systemRoleName).
		First(&role, id).
		Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RolesServices) GetRoleByName(name string) (*entities.Roles, error) {
	var role entities.Roles
	if err := r.db.
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		}).
		Where("name = ?", name).
		First(&role).
		Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RolesServices) GetRolePermissions(roleID uint) ([]entities.Permissions, error) {
	systemRoleName := os.Getenv("SYS_ROLE_NAME")

	var permissions []entities.Permissions
	if err := r.db.
		Where("role_id = ? AND role_name != ?", roleID, systemRoleName).
		Find(&permissions).
		Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *RolesServices) UpdateRole(id uint, role *entities.UpdateRoles) error {
	systemRoleName := os.Getenv("SYS_ROLE_NAME")
	return r.db.
		Model(&entities.Roles{}).
		Where("id = ? AND name != ?", id, systemRoleName).
		Updates(role).Error
}

func (r *RolesServices) DeleteRole(id uint) error {
	systemRoleName := os.Getenv("SYS_ROLE_NAME")

	return r.db.
		Where("name != ?", systemRoleName).
		Delete(&entities.Roles{}, id).
		Error
}

func (r *RolesServices) CreateSystemRole() error {
	roleName := os.Getenv("SYS_ROLE_NAME")

	// Check if role ID and name are provided
	if roleName == "" {
		log.Fatal("System role ID or name not provided in .env file")
		return errors.New("system role ID or name not provided")
	}

	// Check if role exists
	var role_count int64
	err := r.db.Model(&entities.Roles{}).Where("name = ?", roleName).Count(&role_count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatal("Error checking role existence: ", err)
		return err
	}

	if role_count == 0 {
		// start transaction
		tx := r.db.Begin()
		if tx.Error != nil {
			log.Fatal("Error starting transaction: ", tx.Error)
			return tx.Error
		}

		// create role
		role := entities.Roles{
			Name: roleName,
		}
		if err := tx.Create(&role).Error; err != nil {
			tx.Rollback()
			log.Fatal("Error creating role: ", err)
			return err
		}

		// commit transaction
		if err := tx.Commit().Error; err != nil {
			log.Fatal("Error committing transaction: ", err)
			return err
		}
	}

	return nil
}

func (r *RolesServices) GetSystemRole() (*entities.Roles, error) {
	roleName := os.Getenv("SYS_ROLE_NAME")
	var role entities.Roles
	if err := r.db.
		Where("name = ?", roleName).
		First(&role).
		Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RolesServices) IsSystemRole(roleID uint) bool {
	roleName := os.Getenv("SYS_ROLE_NAME")
	var role entities.Roles
	if err := r.db.
		Where("name = ?", roleName).
		First(&role).
		Error; err != nil {
		return false
	}
	return role.ID == roleID
}
