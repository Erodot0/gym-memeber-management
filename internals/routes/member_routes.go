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
		Services: &memberServices,
	}

	userMiddleware := middlewares.UserMiddlewares{
		UserService: &userServices,
		Http:        &secondary.HttpServices{},
	}

	memberMiddleware := middlewares.MemberMiddlewares{
		Services: &memberServices,
		Parser:   &primary.ErrorHandler{},
		Http:     &secondary.HttpServices{},
	}

	r.App.Post("/api/v1/members", userMiddleware.AuthorizeUser, userMiddleware.OnlyOwner, memberHandler.CreateMember)
	r.App.Get("/api/v1/members", userMiddleware.AuthorizeUser, memberHandler.GetMembers)

	r.App.Get("/api/v1/members/:id", userMiddleware.AuthorizeUser, memberMiddleware.GetMember, memberHandler.GetMemberById)
	r.App.Put("/api/v1/members/:id", userMiddleware.AuthorizeUser, userMiddleware.OnlyOwner, memberMiddleware.GetMember, memberHandler.UpdateMember)
	r.App.Delete("/api/v1/members/:id", userMiddleware.AuthorizeUser, userMiddleware.OnlyOwner, memberMiddleware.GetMember, memberHandler.DeleteMember)

	//Subscription
	r.App.Post("/api/v1/members/:id/subscriptions", userMiddleware.AuthorizeUser, memberMiddleware.GetMember, memberHandler.CreateMemberSubscription)
	r.App.Get("/api/v1/members/:id/subscriptions", userMiddleware.AuthorizeUser, memberMiddleware.GetMember, memberHandler.GetMemberSubscriptions)
	r.App.Get("/api/v1/members/:id/subscriptions/:sub_id", userMiddleware.AuthorizeUser, memberMiddleware.GetMember, memberHandler.GetMemberSubscriptionById)
	r.App.Put("/api/v1/members/:id/subscriptions/:sub_id", userMiddleware.AuthorizeUser, memberMiddleware.GetMember, memberHandler.UpdateMemberSubscription)
	r.App.Delete("/api/v1/members/:id/subscriptions/:sub_id", userMiddleware.AuthorizeUser, memberMiddleware.GetMember, memberHandler.DeleteMemberSubscription)
}
