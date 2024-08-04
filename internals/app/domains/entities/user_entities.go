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
	Role     *Roles `json:"role,omitempty" gorm:"foreignKey:RoleID;references:ID"`
}

type UserLogin struct {
	Email    string `json:"email" gorm:"unique;not null;index"`
	Password string `json:"password" gorm:"not null"`
}

type UpdateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	RoleID  uint   `json:"role_id"`
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

func (u *UserLogin) Validate() error {
	//Check the email
	if u.Email == "" {
		return fmt.Errorf("l'email è obbligatoria")
	}

	//Check the password
	if u.Password == "" {
		return fmt.Errorf("la password è obbligatoria")
	}

	return nil
}

func (u *User) RemovePassword() {
	u.Password = ""
}

func (u *User) NewRefreshToken(c *fiber.Ctx, token string) Session {
	return Session{
		Token:     token,
		UserID:    u.ID,
		IPAddress: c.IP(),
		UserAgent: c.Get("user-agent"),
		Expires:   time.Hour * 24 * 7, // 7 days
	}
}

func (u *User) NewSessionToken(c *fiber.Ctx, token string) Session {
	return Session{
		Token:     token,
		UserID:    u.ID,
		IPAddress: c.IP(),
		UserAgent: c.Get("user-agent"),
		Expires:   time.Hour * 1, // 1 hour
	}
}

func (u *User) NewRefreshCookie(token string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24 * 7), // 7 days
		HTTPOnly: true,
		Secure:   false, // Ensure this is false in development (HTTP)
		SameSite: fiber.CookieSameSiteLaxMode,
		Path:     "/",
	}
}

func (u *User) RemoveRefreshCookie() *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now(),
		HTTPOnly: true,
		Secure:   false, // Ensure this is false in development (HTTP)
		SameSite: "None",
		Path:     "/",
	}
}

func (u *User) NewSessionCookie(token string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(1 * time.Hour), // 1 hours session
		HTTPOnly: true,
		Secure:   false, // Ensure this is false in development (HTTP)
		SameSite: "None",
		Path:     "/",
	}
}

func (u *User) RemoveSessionCookie() *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now(),
		HTTPOnly: true,
		Secure:   false, // Ensure this is false in development (HTTP)
		SameSite: "None",
		Path:     "/",
	}
}

// flag only for front end to know if user is logged in,
// same expires time as refresh token
func (u *User) NewLoginCookie() *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "login_token",
		Value:    "true",
		Expires:  time.Now().Add(time.Hour * 24 * 7), // 7 days
		HTTPOnly: false,
		Secure:   false,
		SameSite: "None",
		Path:     "/",
	}
}

func (u *User) RemoveloginCookie() *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "login_token",
		Value:    "",
		Expires:  time.Now(),
		HTTPOnly: false,
		Secure:   false,
		SameSite: "None",
		Path:     "/",
	}
}
