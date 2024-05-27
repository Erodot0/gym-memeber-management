package ports

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
)

type MemberServices interface {
	// CreateMember creates a new member in the database.
	//
	// Parameters:
	//   - m: the member entity to be created.
	//
	// Return type:
	//   - error
	CreateMember(m *entities.Member) error
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
	//   - m: the member entity to be retrieved, it must have the ID field set to the desired member's ID.
	//
	// Return type:
	//   - error: an error if the retrieval process encounters any issues.
	//
	GetMemberById(m *entities.Member) error
	// DeleteMember deletes a member from the database.
	//		Note: It deletes the member and its associated entities.
	// Parameters:
	//   - m: the member entity to be deleted.
	//
	// Return type:
	//   - error: an error if the deletion process encounters any issues.
	//
	DeleteMember(m *entities.Member) error
}
