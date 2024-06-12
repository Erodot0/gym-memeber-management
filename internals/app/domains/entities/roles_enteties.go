package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type Roles struct {
	gorm.Model
	Name        string        `json:"name" gorm:"unique;not null;index"`
	Users       []User        `json:"users,omitempty" gorm:"foreignKey:RoleID"`
	Permissions []Permissions `json:"permissions,omitempty" gorm:"foreignKey:RoleId"`
}

type UpdateRoles struct {
	Name string `json:"name" gorm:"unique;not null;index"`
}

// Notes:
//   - 0 -> no access
//   - 1 -> access
//   - 2 -> self access (only for it self)
type Permissions struct {
	gorm.Model
	Table  string `json:"table" gorm:"not null;index"`
	RoleId uint   `json:"role_id" gorm:"not null;index"`
	Create uint   `json:"create"`
	Read   uint   `json:"read"`
	Update uint   `json:"update"`
	Delete uint   `json:"delete"`
}

type UpdatePermissions struct {
	Create uint `json:"create"`
	Read   uint `json:"read"`
	Update uint `json:"update"`
	Delete uint `json:"delete"`
}

func (r *Roles) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("inserire un nome")
	}

	return nil
}

func (r *UpdateRoles) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("inserire un nome")
	}

	return nil
}

func (r *Permissions) Validate() error {
	if r.Table == "" {
		return fmt.Errorf("inserire una tabella")
	}

	if r.RoleId == 0 {
		return fmt.Errorf("inserire un ruolo")
	}

	if r.Create > 1 || r.Update > 2 || r.Read > 2 || r.Delete > 2 {
		return fmt.Errorf("inserire un permesso valido")
	}

	return nil
}
