package ports

import "github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"

type RolesServices interface {

	// CreateRole creates a new role in the system.
	//
	// Parameters:
	// - role: A pointer to the Roles struct representing the new role to be created.
	//
	// Returns:
	// - error: An error object if there was an issue creating the role, otherwise nil.
	//
	CreateRole(role *entities.Roles) error

	// GetAllRoles retrieves all roles from the system.
	//
	// It returns a slice of entities.Roles and an error if any occurred.
	//
	GetAllRoles() ([]entities.Roles, error)

	// GetRole retrieves a role from the system by its ID.
	//
	// Parameters:
	// - id: The ID of the role to retrieve.
	//
	// Returns:
	// - *entities.Roles: A pointer to the Roles struct representing the retrieved role, or nil if not found.
	// - error: An error object if there was an issue retrieving the role, otherwise nil.
	//
	GetRole(id uint) (*entities.Roles, error)

	// GetRoleByName retrieves a role from the system by its name.
	//
	// Parameters:
	// - name: The name of the role to retrieve.
	//
	// Returns:
	// - *entities.Roles: A pointer to the Roles struct representing the retrieved role, or nil if not found.
	// - error: An error object if there was an issue retrieving the role, otherwise nil.
	//
	GetRoleByName(name string) (*entities.Roles, error)

	// GetRolePermissions retrieves the permissions of a role from the system by its ID.
	//
	// Parameters:
	// - roleID: The ID of the role to retrieve the permissions for.
	//
	// Returns:
	// - []entities.Permissions: A slice of entities.Permissions representing the permissions of the role.
	// - error: An error object if there was an issue retrieving the permissions, otherwise nil.
	//
	GetRolePermissions(roleID uint) ([]entities.Permissions, error)

	// UpdateRole updates a role in the system by its ID.
	//
	// Parameters:
	// - id: The ID of the role to be updated.
	// - role: A pointer to the UpdateRoles struct representing the updated role.
	//
	// Returns:
	// - error: An error object if there was an issue updating the role, otherwise nil.
	//
	UpdateRole(id uint, role *entities.UpdateRoles) error

	// DeleteRole deletes a role from the system by its ID.
	//
	// Parameters:
	// - id: The ID of the role to be deleted.
	//
	// Returns:
	// - error: An error object if there was an issue deleting the role, otherwise nil.
	//
	DeleteRole(id uint) error

	// CreateSystemRole creates a new system role in the system.
	//
	// Returns:
	// - error: An error object if there was an issue creating the system role, otherwise nil.
	//
	CreateSystemRole() error

	// GetSystemRole retrieves the system role from the system.
	//
	// Returns:
	// - *entities.Roles: A pointer to the Roles struct representing the system role, or nil if not found.
	// - error: An error object if there was an issue retrieving the system role, otherwise nil.
	//
	GetSystemRole() (*entities.Roles, error)

	// IsSystemRole checks if a role is a system role.
	//
	// Parameters:
	// - roleID: The ID of the role to check.
	//
	// Returns:
	// - bool: true if the role is a system role, false otherwise.
	//
	IsSystemRole(roleID uint) bool
}
