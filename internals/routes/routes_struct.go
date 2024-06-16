package routes

import (
	primary "github.com/Erodot0/gym-memeber-management/internals/adapters/primary"
	secondary "github.com/Erodot0/gym-memeber-management/internals/adapters/secondary"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/services"
	"github.com/Erodot0/gym-memeber-management/internals/app/handlers"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/middlewares"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Routes struct {
	app   *fiber.App
	db    *gorm.DB
	cache *redis.Client

	// Middlewares
	userMiddlewares   *middlewares.UserMiddlewares
	memberMiddlewares *middlewares.MemberMiddlewares

	// Handlers
	permissionHandlers *handlers.PermissionsHandler
	memberHandlers     *handlers.MembersHandlers
	userHandlers       *handlers.UserHandlers
	roleHandlers       *handlers.RolesHandlers

	// Routes
	authRoutes      fiber.Router
	publicRoutes    fiber.Router
	protectedRoutes fiber.Router
}

// NewRoutes creates a new Routes struct.
func NewRoutes(app *fiber.App, db *gorm.DB, cache *redis.Client) *Routes {

	// Adapters
	httpAdapters := secondary.NewHttpServices()
	cacheAdapters := secondary.NewCacheServices(cache)
	parserAdapters := primary.NewErrorHandler()

	// Services
	userServices := services.NewUserServices(db, cacheAdapters)
	memberServices := services.NewMemberServices(db)
	rolesServices := services.NewRolesServices(db)
	permissionsServices := services.NewPermissionsService(db)

	// Middlewares
	userMiddlewares := middlewares.NewUserMiddlewares(httpAdapters, userServices, permissionsServices)
	memberMiddlewares := middlewares.NewMemberMiddlewares(httpAdapters, memberServices)

	// Handlers
	userHandlers := handlers.NewUserHandlers(parserAdapters, httpAdapters, userServices, rolesServices)
	memberHandlers := handlers.NewMembersHandlers(parserAdapters, httpAdapters, memberServices)
	rolesHandlers := handlers.NewRolesHandlers(parserAdapters, httpAdapters, rolesServices)
	permissionsHandlers := handlers.NewPermissionsHandler(parserAdapters, httpAdapters, permissionsServices)

	// apis
	api := app.Group("/api")

	// v1
	v1 := api.Group("/v1")

	// Public routes group
	authRoutes := v1.Group("/auth")
	publicApi := v1.Group("/public")
	protectedApi := v1.Group("/protected", userMiddlewares.AuthorizeUser, userMiddlewares.CheckPermissions)

	return &Routes{
		app:   app,
		db:    db,
		cache: cache,

		// Middlewares
		userMiddlewares:   userMiddlewares,
		memberMiddlewares: memberMiddlewares,

		// Handlers
		userHandlers:       userHandlers,
		memberHandlers:     memberHandlers,
		roleHandlers:       rolesHandlers,
		permissionHandlers: permissionsHandlers,

		// Routes
		authRoutes:      authRoutes,
		publicRoutes:    publicApi,
		protectedRoutes: protectedApi,
	}
}
