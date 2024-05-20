package entities

import "time"

type Session struct {
	Token     string        `json:"token" gorm:"unique;not null;index"`
	IPAddress string        `json:"ip_address" gorm:"not null;index"`
	UserAgent string        `json:"user_agent" gorm:"not null;index"`
	UserID    uint          `json:"user_id" gorm:"not null;index"`
	Expires   time.Duration `json:"expires" gorm:"not null"`
}