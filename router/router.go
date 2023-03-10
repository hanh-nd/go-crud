package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"hanhngo.me/m/modules/auth"
	"hanhngo.me/m/modules/permissions"
	roleGroups "hanhngo.me/m/modules/role-groups"
	"hanhngo.me/m/modules/roles"
	userGroups "hanhngo.me/m/modules/user-groups"
	"hanhngo.me/m/modules/users"
	"hanhngo.me/m/plugins/jwt"
)

var (
	userService       = users.NewUserService()
	userHandler       = users.NewUserHandler(userService)
	jwtService        = jwt.NewJwtService()
	authService       = auth.NewAuthService(userService, jwtService)
	authHandler       = auth.NewAuthHandler(authService, userService)
	permissionService = permissions.NewPermissionService()
	permissionHandler = permissions.NewPermissionHandler(permissionService)
	roleService       = roles.NewRoleService(permissionService)
	roleHandler       = roles.NewRoleHandler(roleService)
	roleGroupService  = roleGroups.NewRoleGroupService()
	roleGroupHandler  = roleGroups.NewRoleGroupHandler(roleGroupService)
	userGroupService  = userGroups.NewUserGroupService()
	userGroupHandler  = userGroups.NewUserGroupHandler(userGroupService)
)

func SetupRouters(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	api := app.Group("/api", logger.New())

	authRoutes := app.Group("/", logger.New())
	authRoutes.Post("/register", authHandler.Register)
	authRoutes.Post("/login", authHandler.Login)

	userRoutes := api.Group("/users", logger.New())
	userRoutes.Get("/", userHandler.GetUserList)
	userRoutes.Get("/:id", userHandler.GetUserById)
	userRoutes.Patch("/:id", userHandler.UpdateUserProfile)
	userRoutes.Delete("/:id", userHandler.DeleteUser)

	roleRoutes := api.Group("/roles", logger.New())
	roleRoutes.Post("/", roleHandler.CreateRole)
	roleRoutes.Get("/", roleHandler.GetRoleList)
	roleRoutes.Get("/:id", roleHandler.GetRoleById)
	roleRoutes.Patch("/:id", roleHandler.UpdateRole)
	roleRoutes.Patch("/:id/change-permissions", roleHandler.UpdateRolePermissions)
	roleRoutes.Delete("/:id", roleHandler.DeleteRole)

	permissionRoutes := api.Group("/permissions", logger.New())
	permissionRoutes.Post("/", permissionHandler.CreatePermission)
	permissionRoutes.Get("/", permissionHandler.GetPermissionList)
	permissionRoutes.Get("/:id", permissionHandler.GetPermissionById)
	permissionRoutes.Patch("/:id", permissionHandler.UpdatePermission)
	permissionRoutes.Delete("/:id", permissionHandler.DeletePermission)

	roleGroupRoutes := api.Group("/role-groups", logger.New())
	roleGroupRoutes.Post("/", roleGroupHandler.CreateRoleGroup)
	roleGroupRoutes.Get("/", roleGroupHandler.GetRoleGroupList)
	roleGroupRoutes.Get("/:id", roleGroupHandler.GetRoleGroupById)
	roleGroupRoutes.Patch("/:id", roleGroupHandler.UpdateRoleGroup)
	roleGroupRoutes.Delete("/:id", roleGroupHandler.DeleteRoleGroup)

	userGroupRoutes := api.Group("/user-groups", logger.New())
	userGroupRoutes.Post("/", userGroupHandler.CreateUserGroup)
	userGroupRoutes.Get("/", userGroupHandler.GetUserGroupList)
	userGroupRoutes.Get("/:id", userGroupHandler.GetUserGroupById)
	userGroupRoutes.Patch("/:id", userGroupHandler.UpdateUserGroup)
	userGroupRoutes.Delete("/:id", userGroupHandler.DeleteUserGroup)
}
