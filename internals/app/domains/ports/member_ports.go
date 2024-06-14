package ports

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
)

type MemberServicesPort interface {

	// CreateMember creates a new member in the database.
	//
	// Parameters:
	//   - m: the member entity to be created.
	//
	// Return type:
	//   - error
	CreateMember(m *entities.Member) error

	// UpdateMember updates a member in the database.
	// 		Note: It updates the member only and not its associated entities.
	//
	// Parameters:
	//   - id: the ID of the member to be updated.
	//   - m: the member entity to be updated.
	//
	// Return type:
	//   - error: an error if the update process encounters any issues.
	//
	UpdateMember(id uint, m *entities.UpdateMember) error

	// GetAllMembers retrieves all members from the database.
	// 		Note: all members are returned regardless of their subscription status.
	// 		Note: only active subscriptions is returned for each member.
	// Return type:
	//   - []entities.Member: a slice of Member entities representing all members.
	//   - error: an error if the retrieval process encounters any issues.
	GetAllMembers() ([]entities.Member, error)

	// GetMemberById retrieves a member from the database by their ID.
	// 		Note: only active subscriptions is returned for the member.
	// Parameters:
	//   - id: the ID of the member to be retrieved.
	//
	// Return type:
	//   - entities.Member: the member entity representing the member with the given ID.
	//   - error: an error if the retrieval process encounters any issues.
	//
	GetMemberById(id uint) (entities.Member, error)

	// DeleteMember deletes a member from the database.
	//		Note: It deletes the member and its associated entities.
	// Parameters:
	//   - id: the ID of the member to be deleted.
	//
	// Return type:
	//   - error: an error if the deletion process encounters any issues.
	//
	DeleteMember(id uint) error

	// CreateSubscription creates a new subscription for a given member ID.
	//
	// Parameters:
	// - user_id: the ID of the member.
	// - subscription: the subscription entity to be created.
	//
	// Return type:
	// - error: an error if the creation process encounters any issues.
	//
	CreateMemberSubscription(user_id uint, subscription *entities.Subscription) error

	// GetAllSubscriptions retrieves all subscriptions for a given member ID.
	//
	// Parameters:
	// - id: the ID of the member.
	//
	// Return type:
	// - []entities.Subscription: a slice of Subscription entities representing all subscriptions.
	// - error: an error if the retrieval process encounters any issues.
	//
	GetAllSubscriptions(id uint) ([]entities.Subscription, error)

	// GetMembersBySubscription retrieves all subscriptions for a given member ID and subscription ID.
	//
	// Parameters:
	// - id: the ID of the member.
	// - sub_id: the ID of the subscription.
	//
	// Return type:
	// - []entities.Subscription: a slice of Subscription entities representing all subscriptions.
	// - error: an error if the retrieval process encounters any issues.
	//
	GetSubscriptionById(id uint, sub_id uint) ([]entities.Subscription, error)

	// UpdateSubscription updates a subscription for a given user and subscription ID.
	//
	// Parameters:
	// - user_id: the ID of the user.
	// - sub_id: the ID of the subscription.
	// - subscription: a pointer to an entities.UpdateSubscription struct containing the new subscription details.
	//
	// Return type:
	// - []entities.Subscription: a slice of entities.Subscription representing the updated subscriptions.
	// - error: an error if the update process encounters any issues.
	UpdateSubscription(user_id uint, sub_id uint, subscription *entities.UpdateSubscription) ([]entities.Subscription, error)

	// DeleteSubscription deletes a subscription for a given user and subscription ID.
	//
	// Parameters:
	// - user_id: the ID of the user.
	// - sub_id: the ID of the subscription.
	//
	// Return type:
	// - error: an error if the deletion process encounters any issues.
	//
	DeleteSubscription(user_id uint, sub_id uint) error
}