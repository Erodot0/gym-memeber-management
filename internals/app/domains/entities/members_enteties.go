package entities

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name         string         `json:"name" gorm:"not null,required"`
	Surname      string         `json:"surname" gorm:"not null,required"`
	Gender       string         `json:"gender"`
	DateOfBirth  time.Time      `json:"date_of_birth" gorm:"not null,required"`
	Contacts     *Contacts      `json:"contacts" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;"`
	Address      *Address       `json:"address" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;"`
	Subscription []Subscription `json:"subscription" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type Contacts struct {
	ID      uint `json:"ID" gorm:"primaryKey;autoIncrement;unique;not null"`
	Deleted gorm.DeletedAt
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type Address struct {
	ID      uint `json:"ID" gorm:"primaryKey;autoIncrement;unique;not null"`
	Deleted gorm.DeletedAt
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

type Subscription struct {
	ID        uint `json:"ID" gorm:"primaryKey;autoIncrement;unique;not null"`
	Deleted   gorm.DeletedAt
	UserID    uint      `json:"user_id"`
	Type      string    `json:"type"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	IsActive  *bool     `json:"is_active"`
	Price     float32   `json:"price"`
}

func (s *Subscription) AddEndDate() {
	if s.Type == "mensile" {
		s.EndDate = time.Now().AddDate(0, 1, 0)
	}

	if s.Type == "trimestrale" {
		s.EndDate = time.Now().AddDate(0, 3, 0)
	}

	if s.Type == "semestrale" {
		s.EndDate = time.Now().AddDate(0, 6, 0)
	}

	if s.Type == "annuale" {
		s.EndDate = time.Now().AddDate(1, 0, 0)
	}

	// Activate subscription
	*s.IsActive = true
}

func (m *Member) Validate() error {
	if m.Name == "" ||
		m.Surname == "" ||
		m.Gender == "" ||
		m.DateOfBirth.IsZero() ||
		m.Contacts == nil ||
		m.Address == nil ||
		len(m.Subscription) == 0 {
		return fmt.Errorf("compilare i campi obbligatori")
	}

	if err := m.Contacts.Validate(); err != nil {
		return err
	}

	if err := m.Address.Validate(); err != nil {
		return err
	}

	if err := m.Subscription[0].Validate(); err != nil {
		return err
	}

	return nil
}

func (c *Contacts) Validate() error {
	if c.Phone == "" {
		return fmt.Errorf("compilare il numero di telefono")
	}
	return nil
}

func (a *Address) Validate() error {
	if a.Country == "" || a.City == "" || a.Street == "" {
		return fmt.Errorf("l'indirizzo deve avere tutti i campi obbligatori")
	}
	return nil
}

func (s *Subscription) Validate() error {
	validTypes := map[string]bool{
		"mensile":     true,
		"trimestrale": true,
		"semestrale":  true,
		"annuale":     true,
	}

	if !validTypes[s.Type] || s.StartDate.IsZero() || s.Price == 0 || s.Price < 0 || s.IsActive == nil {
		return fmt.Errorf("compilare i campi d'abbonamento correttamente")
	}
	return nil
}
