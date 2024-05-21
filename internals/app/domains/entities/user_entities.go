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

func (u *User) Validate(phase string) error {
	//Check the role
	if (u.Role != "owner" && u.Role != "admin" && phase == "register") {
		return fmt.Errorf("invalid role")
	}

	//Check the email
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}

	//Check the password
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}

	return nil
}

func (u *User) RemovePassword() {
	u.Password = ""
}
