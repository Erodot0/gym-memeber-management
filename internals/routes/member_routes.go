package routes

func (r *Routes) RegisterMemberRoutes() {
	r.protectedRoutes.Post("/members", r.memberHandlers.CreateMember)
	r.protectedRoutes.Get("/members", r.memberHandlers.GetMembers)

	r.protectedRoutes.Get("/members/:id", r.memberMiddlewares.GetMember, r.memberHandlers.GetMemberById)
	r.protectedRoutes.Put("/members/:id", r.memberMiddlewares.GetMember, r.memberHandlers.UpdateMember)
	r.protectedRoutes.Delete("/members/:id", r.memberMiddlewares.GetMember, r.memberHandlers.DeleteMember)

	//Subscription
	r.protectedRoutes.Post("/members/:id/subscriptions", r.memberMiddlewares.GetMember, r.memberHandlers.CreateMemberSubscription)
	r.protectedRoutes.Get("/members/:id/subscriptions", r.memberMiddlewares.GetMember, r.memberHandlers.GetMemberSubscriptions)
	r.protectedRoutes.Get("/members/:id/subscriptions/:sub_id", r.memberMiddlewares.GetMember, r.memberHandlers.GetMemberSubscriptionById)
	r.protectedRoutes.Put("/members/:id/subscriptions/:sub_id", r.memberMiddlewares.GetMember, r.memberHandlers.UpdateMemberSubscription)
	r.protectedRoutes.Delete("/members/:id/subscriptions/:sub_id", r.memberMiddlewares.GetMember, r.memberHandlers.DeleteMemberSubscription)
}
