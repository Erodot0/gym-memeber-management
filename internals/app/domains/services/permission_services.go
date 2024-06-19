package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PermissionsService struct {
	db            *gorm.DB
	rolesServices ports.RolesServices
}

func NewPermissionsService(db *gorm.DB, rolesServices ports.RolesServices) *PermissionsService {
	return &PermissionsService{
		db:            db,
		rolesServices: rolesServices,
	}
}

func (p *PermissionsService) ValidateNewPermission(perm *entities.Permissions) error {
	// Check if the permission already exists
	exists, err := p.CheckPermissionExists(perm.TableName, perm.RoleId)
	if err != nil {
		return fmt.Errorf("errore nel controllo dell'esistenza del permesso")
	}
	if exists {
		return fmt.Errorf("i permessi per questa tabella e ruolo sono già presenti")
	}
	// Check if the role exists
	if perm.RoleId == 0 {
		return fmt.Errorf("il ruolo è obbligatorio")
	}

	// Check role in the database
	role := &entities.Roles{}
	if err := p.db.First(role, perm.RoleId).Error; err != nil {
		return fmt.Errorf("il ruolo non è valido")
	}

	// Check if the table_name exists
	if perm.TableName == "" {
		return fmt.Errorf("la tabella è obbligatoria")
	}

	// Check if the table exists
	if !p.db.Migrator().HasTable(perm.TableName) {
		return fmt.Errorf("la tabella non è valida")
	}
	// Check if the action is valid
	if (perm.Create != nil && *perm.Create > uint(1)) ||
		(perm.Update != nil && *perm.Update > uint(2)) ||
		(perm.Read != nil && *perm.Read > uint(2)) ||
		(perm.Delete != nil && *perm.Delete > uint(2)) {
		return fmt.Errorf("assegnare i permessi correttamente")
	}

	return nil
}

func (p *PermissionsService) ValidateUpdatePermission(perm *entities.UpdatePermissions) error {
	// Check if the action is valid
	if (perm.Create != nil && *perm.Create > uint(1)) ||
		(perm.Update != nil && *perm.Update > uint(2)) ||
		(perm.Read != nil && *perm.Read > uint(2)) ||
		(perm.Delete != nil && *perm.Delete > uint(2)) {
		return fmt.Errorf("assegnare i permessi correttamente")
	}

	return nil
}

func (p *PermissionsService) CreatePermission(perm *entities.Permissions) error {
	return p.db.Create(perm).Error
}

func (p *PermissionsService) GetPermission(id uint) (*entities.Permissions, error) {
	perm := &entities.Permissions{}
	return perm, p.db.First(perm, id).Error
}

func (p *PermissionsService) GetAllPermissions() ([]entities.Permissions, error) {
	perms := []entities.Permissions{}
	return perms, p.db.Find(&perms).Error
}

func (p *PermissionsService) GetPermissionsByRole(roleId uint) ([]entities.Permissions, error) {
	perms := []entities.Permissions{}
	return perms, p.db.Where("role_id = ?", roleId).Find(&perms).Error
}

func (p *PermissionsService) GetPermissionsByTable(table_name string) ([]entities.Permissions, error) {
	perms := []entities.Permissions{}
	return perms, p.db.Where("table_name = ?", table_name).Find(&perms).Error
}

func (p *PermissionsService) UpdatePermission(id uint, perm *entities.UpdatePermissions) (*entities.Permissions, error) {
	if err := p.db.
		Model(&entities.Permissions{}).
		Where("id = ?", id).
		Updates(perm).Error; err != nil {
		return nil, err
	}

	return p.GetPermission(id)
}

func (p *PermissionsService) DeletePermission(id uint) error {
	return p.db.Delete(&entities.Permissions{}, id).Error
}

func (p *PermissionsService) HasPermission(table_name string, roleId uint, action string) (uint, error) {
	var perm uint

	if err := p.db.
		Model(&entities.Permissions{}).
		Select(action).
		Where("table_name = ? AND role_id = ?", table_name, roleId).
		Scan(&perm).Error; err != nil {
		return 0, err
	}

	return perm, nil
}

func (p *PermissionsService) CheckPermissionExists(table_name string, roleId uint) (bool, error) {
	var count int64
	err := p.db.Model(&entities.Permissions{}).
		Where("table_name = ? AND role_id = ?", table_name, roleId).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (p *PermissionsService) GetTableList() ([]string, error) {
	return p.db.Migrator().GetTables()
}

func (p *PermissionsService) GetRequestedActionAndTable(c *fiber.Ctx) (action string, table string) {
	method := c.Method()
	endpoint := c.Path()

	// Get the requested action
	switch method {
	case "POST":
		action = "create"
	case "GET":
		action = "read"
	case "PUT":
		action = "update"
	case "DELETE":
		action = "delete"
	default:
		action = ""
	}

	segments := strings.Split(endpoint, "/")[4:]

	// Regular expression to match only segments with letters
	regex := regexp.MustCompile(`^[a-zA-Z]+$`)

	var result []string
	for _, segment := range segments {
		if regex.MatchString(segment) {
			result = append(result, segment)
		}
	}

	// Clean the endpoint
	table = result[len(result)-1]
	return action, table
}

func (p *PermissionsService) CreateSystemPermissions() error {
	roleName := os.Getenv("SYS_ROLE_NAME")
	if roleName == "" {
		log.Fatal("SYS_ROLE_NAME not found in .env file")
		return errors.New("SYS_ROLE_NAME not found in .env file")
	}

	// Get the system role
	role, err := p.rolesServices.GetRoleByName(roleName)
	if err != nil {
		log.Fatal("Error getting system role: ", err)
		return err
	}

	fmt.Printf("Role ID: %d\n", role.ID)

	// Check if permissions exists and how many
	var permissions_count int64
	err = p.db.Model(&entities.Permissions{}).Where("role_id = ?", role.ID).Count(&permissions_count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Fatal("Error checking if Permissions exists: ", err)
		return err
	}

	fmt.Println(permissions_count)

	// Get the list of tables
	tables, err := p.GetTableList()
	if err != nil {
		log.Fatal("Error getting tables: ", err)
		return err
	}

	// Check the lenght between the tables and the permissions
	if permissions_count < int64(len(tables)) {
		// start transaction
		tx := p.db.Begin()
		if tx.Error != nil {
			log.Fatal("Error starting transaction: ", tx.Error)
			return tx.Error
		}

		// create permissions
		perm := new(uint)
		*perm = uint(1)

		for _, table := range tables {
			// Check if the permission already exists
			exists, err := p.CheckPermissionExists(table, role.ID)
			if err != nil {
				tx.Rollback()
				log.Fatal("Error checking if Permissions exists: ", err)
				return err
			}

			// If the permission already exists, skip it
			if exists {
				continue
			}
			permission := entities.Permissions{
				TableName: table,
				RoleId:    role.ID,
				Create:    perm,
				Read:      perm,
				Update:    perm,
				Delete:    perm,
			}

			if err := tx.Create(&permission).Error; err != nil {
				tx.Rollback()
				log.Fatal("Error creating permission: ", err)
				return err
			}
		}

		// commit transaction
		if err := tx.Commit().Error; err != nil {
			log.Fatal("Error committing transaction: ", err)
			return err
		}
	}

	return nil
}
