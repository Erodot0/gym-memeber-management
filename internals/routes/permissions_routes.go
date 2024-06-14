package routes

func (r *Routes) RegisterPermissionsRoutes() {
	r.protectedRoutes.Post("/permissions", r.permissionHandlers.CreatePermission)
	r.protectedRoutes.Put("/permissions/:perm_id", r.permissionHandlers.UpdatePermission)
	r.protectedRoutes.Get("/permissions/:perm_id", r.permissionHandlers.GetPermission)
	r.protectedRoutes.Get("/permissions", r.permissionHandlers.GetPermissions)
	r.protectedRoutes.Delete("/permissions/:perm_id", r.permissionHandlers.DeletePermission)
}
