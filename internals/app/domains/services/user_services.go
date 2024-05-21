package services

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserServices struct{}

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
