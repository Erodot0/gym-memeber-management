package routes

import (
	primary "github.com/Erodot0/gym-memeber-management/internals/adapters/primary"
	secondary "github.com/Erodot0/gym-memeber-management/internals/adapters/secondary"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/services"
	"github.com/Erodot0/gym-memeber-management/internals/app/handlers"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/middlewares"
)

func (r *Routes) RegisterPermissionsRoutes() {
	permissionsService := services.NewPermissionsService(r.DB)
	permissionsHandler := handlers.NewPermissionsHandler(
		permissionsService,
		&secondary.HttpServices{},
		&primary.ErrorHandler{},
	)

	userServices := services.UserServices{
		DB: r.DB,
		Cache: &secondary.CacheServices{
			CacheClient: r.Cache,
		},
	}

	userMiddleware := middlewares.UserMiddlewares{
		UserService: &userServices,
		Http:        &secondary.HttpServices{},
	}

	r.App.Post("/api/v1/permissions", userMiddleware.AuthorizeUser, permissionsHandler.CreatePermission)
	r.App.Put("/api/v1/permissions/:perm_id", userMiddleware.AuthorizeUser, permissionsHandler.UpdatePermission)
	r.App.Get("/api/v1/permissions/:perm_id", userMiddleware.AuthorizeUser, permissionsHandler.GetPermission)
	r.App.Get("/api/v1/permissions", userMiddleware.AuthorizeUser, permissionsHandler.GetPermissions)
	r.App.Delete("/api/v1/permissions/:perm_id", userMiddleware.AuthorizeUser, permissionsHandler.DeletePermission)
}
