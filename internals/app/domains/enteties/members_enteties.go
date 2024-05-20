package enteties

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name         string         `json:"name" gorm:"not null,required"`
	Surname      string         `json:"surname" gorm:"not null,required"`
	Gender       string         `json:"gender"`
	DateOfBirth  time.Time      `json:"date_of_birth" gorm:"not null,required"`
	Contacts     *Contacts      `json:"contacts,omitempty" gorm:"foreignKey:ID"`
	Address      *Address       `json:"address,omitempty" gorm:"foreignKey:ID"`
	Subscription []Subscription `json:"subscription,omitempty" gorm:"foreignKey:ID"`
}

type Contacts struct {
	ID    uint   `json:"-" gorm:"primaryKey;autoIncrement;unique;not null"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Address struct {
	ID      uint   `json:"-" gorm:"primaryKey;autoIncrement;unique;not null"`
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Number  string `json:"number"`
}

type Subscription struct {
	ID        uint      `json:"-" gorm:"primaryKey;autoIncrement;unique;not null"`
	Type      string    `json:"type"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	IsActive  *bool     `json:"is_active"`
	Price     int       `json:"price"`
}
