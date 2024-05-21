package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"not null,unique;index"`
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"`
}

func (u *User) RemovePassword() {
	u.Password = ""
}
