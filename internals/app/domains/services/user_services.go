package services

import (
	"log"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServices struct {
	DB *gorm.DB
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
