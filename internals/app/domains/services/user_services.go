package services

import (
	"fmt"
	"log"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServices struct {
	DB    *gorm.DB
	Cache ports.CacheAdapters
}

func (s *UserServices) EcnrypPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("@EcnrypPassword Error hashing password: %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *UserServices) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *UserServices) CreateUser(user *entities.User) error {
	return s.DB.
		Model(user).
		Create(user).
		Error
}

func (s *UserServices) DeleteUser(u *entities.User) error {
	return s.DB.
		Model(u).
		Where("id = ?", u.ID).
		Delete(u).
		Error
}

func (s *UserServices) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	return users, s.DB.
		Model(&users).
		Preload("Role").
		Select("ID", "name", "surname", "email", "created_at", "updated_at", "role_id").
		Find(&users).
		Error
}

func (s *UserServices) GetUserById(u *entities.User) error {
	return s.DB.
		Model(u).
		Where("id = ?", u.ID).
		First(u).
		Error
}

func (s *UserServices) GetUserByEmail(user *entities.User) error {
	return s.DB.
		Model(user).
		Where("email = ?", user.Email).
		First(user).
		Error
}

func (s *UserServices) UpdateUser(id uint, u *entities.UpdateUser) (*entities.User, error) {
	if err := s.DB.Model(&entities.User{}).Where("id = ?", id).Updates(u).Error; err != nil {
		return nil, err
	}

	user := &entities.User{}

	if err := s.DB.
		Model(user).
		Where("id = ?", id).
		Preload("Role").
		Select("ID", "name", "surname", "email", "created_at", "updated_at", "role_id").
		First(user).
		Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServices) SetSession(c *fiber.Ctx, user *entities.User) error {
	//Generate random token
	token, err := utils.GenerateRandomToken(32)
	if err != nil {
		return err
	}

	//Create session
	session := user.NewSession(c, token)

	//Set session in cache
	if err := s.Cache.SetCache(&session); err != nil {
		return err
	}

	//Set cookie
	c.Cookie(user.NewAuthCookie(token))

	return nil
}

func (u *UserServices) GetSessionByToken(token string) (*entities.Session, error) {
	// Create session
	session := &entities.Session{
		Token: token,
	}

	// Get all keys for the token
	keys, err := u.Cache.GetCacheKeys(session)
	if err != nil {
		log.Printf("@GetUserSessionByToken: Error getting keys: %v", err)
		return nil, err
	}

	// Check if 1 key is found, 1 to void multiple keys
	if len(keys) != 1 {
		log.Printf("@GetUserSessionByToken: No key or multiple keys found for pattern")
		return nil, fmt.Errorf("no key or multiple keys found for pattern")
	}

	// Get the session from Redis with key
	if err := u.Cache.GetCacheFromKey(keys[0], session); err != nil {
		log.Printf("@GetUserSessionByToken: Error getting session: %v", err)
		return nil, err
	}

	return session, nil
}

func (u *UserServices) DeleteSession(c *fiber.Ctx, id uint) error {
	// Get authorization token and create session
	token := c.Cookies("Authorization")
	session := entities.Session{
		Token:  token,
		UserID: id,
	}

	// Remove the session from Redis
	if err := u.Cache.DelCache(&session); err != nil {
		log.Printf("@DeleteSession: Error removing session: %v", err)
		return err
	}

	// Clear the cookie
	c.ClearCookie("Authorization")
	return nil
}

func (u *UserServices) DeleteAllSessions(c *fiber.Ctx, id uint) error {
	// Create session
	session := entities.Session{
		UserID: id,
	}

	// Delete the sessions from Redis
	if err := u.Cache.DelCacheMultiple(&session); err != nil {
		log.Printf("@DeleteAllSessions: Error removing session: %v", err)
		return err
	}

	return nil
}
