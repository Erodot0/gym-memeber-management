package routes

import (
	primary "github.com/Erodot0/gym-memeber-management/internals/adapters/primary"
	secondary "github.com/Erodot0/gym-memeber-management/internals/adapters/secondary"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/services"
	"github.com/Erodot0/gym-memeber-management/internals/app/handlers"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/middlewares"
)

func (r *Routes) RegisterRolesRoutes() {
	permissionServices := services.NewPermissionsService(r.DB)
	permissionsHandler := handlers.NewPermissionsHandler(permissionServices, &secondary.HttpServices{}, &primary.ErrorHandler{})
	userServices := services.UserServices{
		DB: r.DB,
		Cache: &secondary.CacheServices{
			CacheClient: r.Cache,
		},
	}

	rolesServices := services.RolesServices{
		DB: r.DB,
	}

	rolesHandler := handlers.RolesHandlers{
		Parser:       &primary.ErrorHandler{},
		Http:         &secondary.HttpServices{},
		RolesService: &rolesServices,
	}

	userMiddleware := middlewares.UserMiddlewares{
		UserService: &userServices,
		Http:        &secondary.HttpServices{},
	}

	r.App.Post("/api/v1/roles", userMiddleware.AuthorizeUser, rolesHandler.CreateRole)
	r.App.Get("/api/v1/roles", userMiddleware.AuthorizeUser, rolesHandler.GetAllRoles)
	r.App.Get("/api/v1/roles/:id", userMiddleware.AuthorizeUser, rolesHandler.GetRole)
	r.App.Put("/api/v1/roles/:id", userMiddleware.AuthorizeUser, rolesHandler.UpdateRole)
	r.App.Delete("/api/v1/roles/:id", userMiddleware.AuthorizeUser, rolesHandler.DeleteRole)

	r.App.Post("/api/v1/roles/:id/permissions", userMiddleware.AuthorizeUser, permissionsHandler.CreatePermission)
	r.App.Get("/api/v1/roles/:id/permissions", userMiddleware.AuthorizeUser, rolesHandler.GerRolePermissions)
	r.App.Put("/api/v1/roles/permissions/:perm_id", userMiddleware.AuthorizeUser, permissionsHandler.UpdatePermission)
	r.App.Delete("/api/v1/roles/permissions/:perm_id", userMiddleware.AuthorizeUser, permissionsHandler.DeletePermission)
}
