package entities

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null,unique;index"`
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname" gorm:"not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`
}

func (u *User) Validate() error {
	//Check the role
	if u.Role != "owner" && u.Role != "admin" {
		return fmt.Errorf("invalid role")
	}

	return nil
}

func (u *User) RemovePassword() {
	u.Password = ""
}
