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
