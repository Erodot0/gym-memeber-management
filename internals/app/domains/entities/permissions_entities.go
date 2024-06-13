package entities

import (
	"gorm.io/gorm"
)

// Notes:
//   - 0 -> no access
//   - 1 -> access
//   - 2 -> self access (only for it self)
type Permissions struct {
	gorm.Model
	TableName string `json:"table_name" gorm:"not null;index"`
	RoleId    uint   `json:"role_id" gorm:"not null;index"`
	Create    uint   `json:"create"`
	Read      uint   `json:"read"`
	Update    uint   `json:"update"`
	Delete    uint   `json:"delete"`
}

type UpdatePermissions struct {
	Create uint `json:"create"`
	Read   uint `json:"read"`
	Update uint `json:"update"`
	Delete uint `json:"delete"`
}
