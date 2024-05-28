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

func (m *MemberServices) UpdateMember(member *entities.UpdateMember) error {
	return m.DB.
		Model(entities.Member{}).
		Where("id = ?", member.ID).
		Updates(member).Error
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

func (m *MemberServices) GetMemberById(id uint) (entities.Member, error) {
	var member entities.Member
	if err := m.DB.
		Preload("Contacts").
		Preload("Address").
		Preload("Subscription", "is_active = true").
		First(&member, id).
		Error; err != nil {
		return member, err
	}
	return member, nil
}

func (m *MemberServices) DeleteMember(id uint) error {
	member := new(entities.Member)
	member.ID = id
	return m.DB.
		Select("Contacts", "Address", "Subscription").
		Delete(member).
		Error
}
