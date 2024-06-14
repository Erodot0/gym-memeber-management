package routes

func (r *Routes) RegisterRolesRoutes() {
	r.protectedRoutes.Post("/roles", r.roleHandlers.CreateRole)
	r.protectedRoutes.Get("/roles", r.roleHandlers.GetAllRoles)
	r.protectedRoutes.Get("/roles/:id", r.roleHandlers.GetRole)
	r.protectedRoutes.Put("/roles/:id", r.roleHandlers.UpdateRole)
	r.protectedRoutes.Delete("/roles/:id", r.roleHandlers.DeleteRole)

	r.protectedRoutes.Post("/roles/:id/permissions", r.permissionHandlers.CreatePermission)
	r.protectedRoutes.Get("/roles/:id/permissions", r.roleHandlers.GerRolePermissions)
	r.protectedRoutes.Put("/roles/permissions/:perm_id", r.permissionHandlers.UpdatePermission)
	r.protectedRoutes.Delete("/roles/permissions/:perm_id", r.permissionHandlers.DeletePermission)
}
