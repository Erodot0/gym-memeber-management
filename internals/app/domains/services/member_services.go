package services

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"gorm.io/gorm"
)

type MemberServices struct {
	DB *gorm.DB
}

func (m *MemberServices) CreateMember(member *entities.Member) error {
	tx := m.DB.Begin()
	if err := tx.Create(member).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (m *MemberServices) GetAllMembers() ([]entities.Member, error) {
	var members []entities.Member
	if err := m.DB.
		Preload("Contacts").
		Preload("Address").
		Preload("Subscription", "is_active = true").
		Find(&members).
		Error; err != nil {
		return nil, err
	}
	return members, nil
}
