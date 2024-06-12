package ports

import "github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"

type PermissionsPort interface {

	// ValidateNewPermission checks if a new permission is valid.
	//
	// Parameters:
	//   - p: a pointer to the Permissions entity representing the permission to be validated.
	//
	// Returns:
	//   - error: an error if the permission validation fails, nil otherwise.
	//
	ValidateNewPermission(p *entities.Permissions) error

	// ValidateUpdatePermission checks if an updated permission is valid.
	//
	// Parameters:
	//   - p: a pointer to the UpdatePermissions entity representing the permission to be validated.
	//
	// Returns:
	//   - error: an error if the permission validation fails, nil otherwise.
	//
	ValidateUpdatePermission(p *entities.UpdatePermissions) error

	// CreatePermission creates a new permission in the system.
	//
	// Parameters:
	//   - p: a pointer to the Permissions entity representing the permission to be created.
	//
	// Returns:
	//   - error: an error if the permission creation fails, nil otherwise.
	//
	CreatePermission(p *entities.Permissions) error

	// GetPermission retrieves a permission from the system by its ID.
	//
	// Parameters:
	//   - id: the ID of the permission to retrieve.
	//
	// Returns:
	//   - *entities.Permissions: a pointer to the Permissions entity representing the retrieved permission, or nil if not found.
	//   - error: an error if the permission retrieval fails, nil otherwise.
	//
	GetPermission(id uint) (*entities.Permissions, error)

	// GetAllPermissions retrieves all permissions from the system.
	//
	// Returns:
	//   - []entities.Permissions: a slice of Permissions entities representing all permissions in the system.
	//   - error: an error if the permission retrieval fails, nil otherwise.
	//
	GetAllPermissions() ([]entities.Permissions, error)

	// GetPermissionsByRole retrieves all permissions for a specific role from the system.
	//
	// Parameters:
	//   - roleId: the ID of the role to retrieve permissions for.
	//
	// Returns:
	//   - []entities.Permissions: a slice of Permissions entities representing all permissions for the specified role.
	//   - error: an error if the permission retrieval fails, nil otherwise.
	//
	GetPermissionsByRole(roleId uint) ([]entities.Permissions, error)

	// GetPermissionsByTable retrieves all permissions for a specific table from the system.
	//
	// Parameters:
	//   - table: the name of the table to retrieve permissions for.
	//
	// Returns:
	//   - []entities.Permissions: a slice of Permissions entities representing all permissions for the specified table.
	//   - error: an error if the permission retrieval fails, nil otherwise.
	//
	GetPermissionsByTable(table string) ([]entities.Permissions, error)

	// UpdatePermission updates a permission in the system by its ID.
	//
	// Parameters:
	//   - id: the ID of the permission to be updated.
	//   - p: a pointer to the UpdatePermissions entity representing the updated permission.
	//
	// Returns:
	//   - *entities.Permissions: a pointer to the Permissions entity representing the updated permission, or nil if not found.
	//   - error: an error if the permission update fails, nil otherwise.
	//
	UpdatePermission(id uint, p *entities.UpdatePermissions) (*entities.Permissions, error)

	// DeletePermission deletes a permission from the system by its ID.
	//
	// Parameters:
	//   - id: the ID of the permission to be deleted.
	//
	// Returns:
	//   - error: an error if the permission deletion fails, nil otherwise.
	//
	DeletePermission(id uint) error

	// HasPermission checks if a permission exists for a specific role and table.
	//
	// Parameters:
	//   - table: the name of the table to check permissions for.
	//   - roleId: the ID of the role to check permissions for.
	//   - action: the action to check permissions for (create, read, update, delete).
	//
	// Returns:
	//   - bool: true if the permission exists, false otherwise.
	//   - error: an error if the permission check fails, nil otherwise.
	//
	HasPermission(table string, roleId uint, action string) (bool, error)

	// CheckPermissionExists checks if a permission exists for a specific role and table.
	//
	// Parameters:
	//   - table: the name of the table to check permissions for.
	//   - roleId: the ID of the role to check permissions for.
	//
	// Returns:
	//   - bool: true if the permission exists, false otherwise.
	//   - error: an error if the permission check fails, nil otherwise.
	//
	CheckPermissionExists(table string, roleId uint) (bool, error)
}
