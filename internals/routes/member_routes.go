package routes

import (
	primary "github.com/Erodot0/gym-memeber-management/internals/adapters/primary"
	secondary "github.com/Erodot0/gym-memeber-management/internals/adapters/secondary"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/services"
	"github.com/Erodot0/gym-memeber-management/internals/app/handlers"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/middlewares"
)

func (r *Routes) RegisterMemberRoutes() {
	userServices := services.UserServices{
		DB: r.DB,
		Cache: &secondary.CacheServices{
			CacheClient: r.Cache,
		},
	}

	memberServices := services.MemberServices{
		DB: r.DB,
	}

	memberHandler := handlers.MembersHandlers{
		Parser:   &primary.ErrorHandler{},
		Http:     &secondary.HttpServices{},
		Services:   &memberServices,
	}

	userMiddleware := middlewares.UserMiddlewares{
		UserService: &userServices,
		Http:        &secondary.HttpServices{},
	}

	r.App.Post("/api/v1/members", userMiddleware.AuthorizeUser, userMiddleware.OnlyOwner, memberHandler.CreateMember)
}
