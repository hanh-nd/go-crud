package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"hanhngo.me/m/modules/auth"
	"hanhngo.me/m/modules/permissions"
	"hanhngo.me/m/modules/roles"
	"hanhngo.me/m/modules/users"
	"hanhngo.me/m/plugins/jwt"
)

var (
	userService       = users.NewUserService()
	userHandler       = users.NewUserHandler(userService)
	jwtService        = jwt.NewJwtService()
	authService       = auth.NewAuthService(userService, jwtService)
	authHandler       = auth.NewAuthHandler(authService, userService)
	roleService       = roles.NewRoleService()
	roleHandler       = roles.NewRoleHandler(roleService)
	permissionService = permissions.NewPermissionService()
	permissionHandler = permissions.NewPermissionHandler(permissionService)
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
	roleRoutes.Delete("/:id", roleHandler.DeleteRole)

	permissionRoutes := api.Group("/permissions", logger.New())
	permissionRoutes.Post("/", permissionHandler.CreatePermission)
	permissionRoutes.Get("/", permissionHandler.GetPermissionList)
	permissionRoutes.Get("/:id", permissionHandler.GetPermissionById)
	permissionRoutes.Patch("/:id", permissionHandler.UpdatePermission)
	permissionRoutes.Delete("/:id", permissionHandler.DeletePermission)
}
