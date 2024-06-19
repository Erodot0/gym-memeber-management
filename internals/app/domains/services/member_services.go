package services

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"gorm.io/gorm"
)

type MemberServices struct {
	db *gorm.DB
}

func NewMemberServices(db *gorm.DB) *MemberServices {
	return &MemberServices{
		db: db,
	}
}

func (m *MemberServices) CreateMember(member *entities.Member) error {
	tx := m.db.Begin()
	if err := tx.Create(member).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (m *MemberServices) UpdateMember(id uint, member *entities.UpdateMember) error {
	return m.db.
		Model(entities.Member{}).
		Where("id = ?", id).
		Updates(member).Error
}

func (m *MemberServices) GetAllMembers() ([]entities.Member, error) {
	var members []entities.Member
	if err := m.db.
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
	if err := m.db.
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
	return m.db.
		Select("Contacts", "Address", "Subscription").
		Delete(member).
		Error
}

func (m *MemberServices) CreateMemberSubscription(user_id uint, subscription *entities.Subscription) error {
	subscription.UserID = user_id
	return m.db.
		Model(entities.Subscription{}).
		Create(subscription).
		Error
}

func (m *MemberServices) GetAllSubscriptions(id uint) ([]entities.Subscription, error) {
	var subscriptions []entities.Subscription
	if err := m.db.
		Model(entities.Subscription{}).
		Where("user_id = ?", id).
		Find(&subscriptions).
		Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (m *MemberServices) GetSubscriptionById(id uint, sub_id uint) ([]entities.Subscription, error) {
	var subscriptions []entities.Subscription
	if err := m.db.
		Model(entities.Subscription{}).
		Where("user_id = ? AND id = ?", id, sub_id).
		First(&subscriptions).
		Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (m *MemberServices) UpdateSubscription(user_id uint, sub_id uint, subscription *entities.UpdateSubscription) ([]entities.Subscription, error) {
	var subscriptions []entities.Subscription
	if err := m.db.
		Model(entities.Subscription{}).
		Where("user_id = ? AND id = ?", user_id, sub_id).
		Updates(subscription).
		Find(&subscriptions).
		Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (m *MemberServices) DeleteSubscription(user_id uint, sub_id uint) error {
	return m.db.
		Where("user_id = ? AND id = ?", user_id, sub_id).
		Delete(&entities.Subscription{}).
		Error
}