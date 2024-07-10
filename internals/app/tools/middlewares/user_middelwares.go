package middlewares

import (
	"log"
	"slices"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
)

type UserMiddlewares struct {
	http                ports.HttpAdapters
	userService         ports.UserServices
	permissionsServices ports.PermissionsServices
}

func NewUserMiddlewares(http ports.HttpAdapters, userService ports.UserServices, permissionsServices ports.PermissionsServices) *UserMiddlewares {
	return &UserMiddlewares{
		http:                http,
		userService:         userService,
		permissionsServices: permissionsServices,
	}
}

func (m *UserMiddlewares) AuthorizeUser(c *fiber.Ctx) error {
	authorization := c.Cookies("Authorization")
	if authorization == "" {
		// Send Unauthorized response
		log.Printf("@AuthorizeUser: Authorization cookie not found")
		return m.http.Forbidden(c)
	}

	// Get session from Redis
	session, err := m.userService.GetSessionByToken(authorization)
	if err != nil {
		log.Printf("@AuthorizeUser: Error getting session: %v", err)
		return m.http.Forbidden(c)
	}

	// Check if it is the same IP and user agent
	if session.IPAddress != c.IP() || session.UserAgent != c.Get("User-Agent") {
		log.Printf("@AuthorizeUser: Session IP/User-Agent mismatch: %v", err)
		return m.http.Forbidden(c)
	}

	// Get user
	user, err := m.userService.GetUserForLogin(session.UserID)
	if err != nil {
		log.Printf("@AuthorizeUser: Error getting user: %v", err)
		return m.http.Forbidden(c)
	}

	// Set session
	utils.SetLocals(c, "session", session)
	utils.SetLocals(c, "user", user)
	utils.SetLocals(c, "role", user.Role)
	return c.Next()
}

func (m *UserMiddlewares) CheckPermissions(c *fiber.Ctx) error {
	// get role
	role := utils.GetLocalRole(c)
	roleId := role.ID

	if roleId == 0 {
		return m.http.Forbidden(c)
	}

	// get requested action
	requestedAction, requestedTable := m.permissionsServices.GetRequestedActionAndTable(c)
	if requestedAction == "" || requestedTable == "" {
		log.Printf("@CheckPermissions: Requested action or table not found")
		return m.http.Forbidden(c)
	}

	// get tables
	tables, err := m.permissionsServices.GetTableList()
	if err != nil {
		log.Printf("@CheckPermissions: Error getting table list: %v", err)
		return m.http.Forbidden(c)
	}

	// Check if the requested table is in the list of tables
	if !slices.Contains(tables, requestedTable) {
		return m.http.Forbidden(c)
	}

	permission, err := m.permissionsServices.HasPermission(requestedTable, roleId, requestedAction)
	if err != nil {
		log.Printf("@CheckPermissions: Error getting permission: %v", err)
		return m.http.Forbidden(c)
	}

	if permission == 0 {
		log.Printf("@CheckPermissions: User doesn't have permission to access this resource: %v, %v", requestedTable, requestedAction)
		return m.http.Forbidden(c)
	}

	// Set permission
	utils.SetLocals(c, "permission", permission)
	return c.Next()
}
