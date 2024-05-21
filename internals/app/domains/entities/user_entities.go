package entities

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
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

func (u *User) NewSession(c *fiber.Ctx, token string) Session {
	return Session{
		Token:  token,
		UserID: u.ID,
		IPAddress: c.IP(),
		UserAgent: c.Get("user-agent"),
		Expires:  10 * time.Hour, // 10 hours session
	}
}

func (u *User) NewAuthCookie(token string) fiber.Cookie {
	return fiber.Cookie{
		Name:     "Authorization",
		Value:    token,
		Expires:  time.Now().Add(10 * time.Hour), // 10 hours session
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteStrictMode,
		Path:     "/",
	}
}
