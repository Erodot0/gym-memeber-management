package services

import (
	"fmt"

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

func (m *MemberServices) GetMemberById(member *entities.Member) error {
	return m.DB.
		Preload("Contacts").
		Preload("Address").
		Preload("Subscription", "is_active = true").
		Where("id = ?", member.ID).
		First(member).
		Error
}

func (m *MemberServices) DeleteMember(member *entities.Member) error {
	fmt.Println("member: ", member.Name)
	return m.DB.
		Select("Contacts", "Address", "Subscription").
		Delete(member).
		Error
}
