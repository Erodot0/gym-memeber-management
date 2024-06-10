package entities

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null;index"`
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname" gorm:"not null"`
	Password string `json:"password,omitempty" gorm:"not null"`
    RoleID   uint   `json:"role_id" gorm:"index"` // Foreign key
    Role     Roles   `json:"role" gorm:"foreignKey:RoleID;references:ID"`
}

func (u *User) Validate() error {
	//Check the email
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}

	//Check the password
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}

	//Check the Role
	if u.RoleID == 0 {
		return fmt.Errorf("role is required")
	}

	return nil
}

func (u *User) ValidateLogin() error {
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

func (u *User) NewSession(c *fiber.Ctx, token string) Session {
	return Session{
		Token:     token,
		UserID:    u.ID,
		IPAddress: c.IP(),
		UserAgent: c.Get("user-agent"),
		Expires:   10 * time.Hour, // 10 hours session
	}
}

func (u *User) NewAuthCookie(token string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "Authorization",
		Value:    token,
		Expires:  time.Now().Add(10 * time.Hour), // 10 hours session
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteStrictMode,
		Path:     "/",
	}
}
