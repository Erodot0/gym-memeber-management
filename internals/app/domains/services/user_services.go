package services

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServices struct {
	db           *gorm.DB
	cache        ports.CacheAdapters
	roleServices ports.RolesServices
}

func NewUserServices(db *gorm.DB, cache ports.CacheAdapters, roleServices ports.RolesServices) *UserServices {
	return &UserServices{
		db:           db,
		cache:        cache,
		roleServices: roleServices,
	}
}

func (s *UserServices) EcnrypPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("@EcnrypPassword Error hashing password: %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *UserServices) ComparePassword(userID uint, password string) error {
	var user entities.User
	if err := s.db.Model(&user).Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (s *UserServices) CreateUser(user *entities.User) error {
	return s.db.
		Model(user).
		Create(user).
		Error
}

func (s *UserServices) DeleteUser(u *entities.User) error {
	return s.db.
		Model(u).
		Where("id = ?", u.ID).
		Delete(u).
		Error
}

func (s *UserServices) GetAllUsers() ([]entities.User, error) {
	systemUserEmail := os.Getenv("SYS_USER_EMAIL")

	var users []entities.User
	return users, s.db.
		Model(&users).
		Preload("Role").
		Omit("password").
		Where("email != ?", systemUserEmail).
		Find(&users).
		Error
}

func (s *UserServices) GetUserById(u *entities.User) error {
	systemUserEmail := os.Getenv("SYS_USER_EMAIL")

	return s.db.
		Model(u).
		Preload("Role").
		Omit("password").
		Where("id = ? AND email != ?", u.ID, systemUserEmail).
		First(u).
		Error
}

func (s *UserServices) GetUserForLogin(id uint) (*entities.User, error) {
	user := &entities.User{}
	return user, s.db.
		Model(user).
		Preload("Role").
		Omit("password").
		Where("id = ?", id).
		First(user).
		Error
}

func (s *UserServices) GetUserByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	if err := s.db.
		Model(user).
		Preload("Role").
		Omit("password").
		Where("email = ?", email).
		First(user).
		Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServices) UpdateUser(id uint, u *entities.UpdateUser) (*entities.User, error) {
	if err := s.db.Model(&entities.User{}).Where("id = ?", id).Updates(u).Error; err != nil {
		return nil, err
	}

	user := &entities.User{}
	if err := s.db.
		Model(user).
		Where("id = ?", id).
		Preload("Role").
		Omit("password").
		First(user).
		Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServices) SetSession(c *fiber.Ctx, user *entities.User) error {
	//Generate random refresh token
	refresh_token, err := utils.GenerateRandomToken(64)
	if err != nil {
		return err
	}
	//Generate random session token
	session_token, err := utils.GenerateRandomToken(32)
	if err != nil {
		return err
	}

	//Create refresh token and set it in cache
	refresh_cache := user.NewRefreshToken(c, refresh_token)
	if err := u.cache.SetCache(&refresh_cache); err != nil {
		return err
	}

	//Create session token and set it in cache
	session_cache := user.NewSessionToken(c, session_token)
	if err := u.cache.SetCache(&session_cache); err != nil {
		return err
	}

	//Set cookies
	c.Cookie(user.NewRefreshCookie(refresh_token))
	c.Cookie(user.NewSessionCookie(session_token))
	c.Cookie(user.NewLoginCookie())

	return nil
}

func (u *UserServices) GetSessionByToken(token string) (*entities.Session, error) {
	// Create session
	session := &entities.Session{
		Token: token,
	}

	// Get all keys for the token
	keys, err := u.cache.GetCacheKeys(session)
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
	if err := u.cache.GetCacheFromKey(keys[0], session); err != nil {
		log.Printf("@GetUserSessionByToken: Error getting session: %v", err)
		return nil, err
	}

	return session, nil
}

func (u *UserServices) GetUserFromSession(c *fiber.Ctx, token string) (*entities.User, error) {
	// Get session from Redis
	session, err := u.GetSessionByToken(token)
	if err != nil {
		log.Printf("@GetUserFromSession: Error getting session: %v", err)
		return nil, err
	}

	// Check if it is the same IP and user agent
	if session.IPAddress != c.IP() || session.UserAgent != c.Get("User-Agent") {
		log.Printf("@GetUserFromSession: Session IP/User-Agent mismatch: %v", err)
		return nil, err
	}

	// Get user
	user, err := u.GetUserForLogin(session.UserID)
	if err != nil {
		log.Printf("@GetUserFromSession: Error getting user: %v", err)
		return nil, err
	}

	return user, nil
}

func (u *UserServices) RefreshUserSession(c *fiber.Ctx, user *entities.User) error {
	//Generate random session token
	session_token, err := utils.GenerateRandomToken(32)
	if err != nil {
		return err
	}

	//Create session token and set it in cache
	session_cache := user.NewSessionToken(c, session_token)
	if err := u.cache.SetCache(&session_cache); err != nil {
		return err
	}

	c.Cookie(user.NewSessionCookie(session_token))
	return nil
}

func (u *UserServices) DeleteSession(c *fiber.Ctx, id uint) error {
	// Get refresh_token token and create session
	token := c.Cookies("refresh_token")
	session := entities.Session{
		Token:  token,
		UserID: id,
	}

	// Remove the session from Redis
	if err := u.cache.DelCache(&session); err != nil {
		log.Printf("@DeleteSession: Error removing session: %v", err)
		return err
	}

	// Clear the cookie
	user := entities.User{}
	c.Cookie(user.RemoveRefreshCookie())
	c.Cookie(user.RemoveSessionCookie())
	c.Cookie(user.RemoveloginCookie())
	return nil
}

func (u *UserServices) DeleteAllSessions(c *fiber.Ctx, id uint) error {
	// Create session
	session := entities.Session{
		UserID: id,
	}

	// Delete the sessions from Redis
	if err := u.cache.DelCacheMultiple(&session); err != nil {
		log.Printf("@DeleteAllSessions: Error removing session: %v", err)
		return err
	}

	return nil
}

func (u *UserServices) CreateSystemUser() error {
	email := os.Getenv("SYS_USER_EMAIL")
	password := os.Getenv("SYS_USER_PWD")
	roleName := os.Getenv("SYS_ROLE_NAME")

	if email == "" || password == "" || roleName == "" {
		log.Fatal("SYS_USER_EMAIL, SYS_USER_PWD or SYS_ROLE_NAME not found in .env file")
		return errors.New("SYS_USER_EMAIL, SYS_USER_PWD or SYS_ROLE_NAME not found in .env file")
	}

	// Get system role
	role, err := u.roleServices.GetRoleByName(roleName)
	if err != nil {
		log.Fatal("Error getting system role: ", err)
		return err
	}

	// check if user exists
	var user_count int64
	err = u.db.Model(&entities.User{}).Where("email = ?", email).Count(&user_count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatal("Error checking user existence: ", err)
		return err
	}

	if user_count == 0 {
		// start transaction
		tx := u.db.Begin()
		if tx.Error != nil {
			log.Fatal("Error starting transaction: ", tx.Error)
			return tx.Error
		}

		// Encrypt password
		hash, err := u.EcnrypPassword(password)
		if err != nil {
			tx.Rollback()
			log.Fatal("Error encrypting password: ", err)
			return err
		}

		// create user
		user := entities.User{
			Name:     "system",
			Surname:  "user",
			Email:    email,
			Password: hash,
			RoleID:   role.ID,
		}
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			log.Fatal("Error creating user: ", err)
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

func (u *UserServices) IsSystemUser(id uint) bool {
	email := os.Getenv("SYS_USER_EMAIL")
	var user entities.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return user.ID == id
}
