package entities

import (
	"fmt"
	"strconv"
	"time"

	"github.com/goccy/go-json"
)

type Session struct {
	Token     string        `json:"token" gorm:"unique;not null;index"`
	IPAddress string        `json:"ip_address" gorm:"not null;index"`
	UserAgent string        `json:"user_agent" gorm:"not null;index"`
	UserID    uint          `json:"user_id" gorm:"not null;index"`
	Expires   time.Duration `json:"expires" gorm:"not null"`
}

func (s *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

func (s *Session) SetCacheKey() string {
	stringId := strconv.FormatUint(uint64(s.UserID), 10)
	return fmt.Sprintf("session:%s:%s", stringId, s.Token)
}

func (s *Session) SetCacheExpiration() time.Duration {
	return s.Expires
}

func (s *Session) GetCacheKey() string {
	stringId := "*"
	token := "*"
	if s.UserID != 0 {
		stringId = strconv.FormatUint(uint64(s.UserID), 10)
	}
	if s.Token != "" {
		token = s.Token
	}
	return fmt.Sprintf("session:%s:%s", stringId, token)
}
