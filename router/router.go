package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"hanhngo.me/m/modules/auth"
	"hanhngo.me/m/modules/users"
	"hanhngo.me/m/plugins/jwt"
)

var (
	userService = users.NewUserService()
	userHandler = users.NewUserHandler(userService)
	jwtService  = jwt.NewJwtService()
	authService = auth.NewAuthService(userService, jwtService)
	authHandler = auth.NewAuthHandler(authService, userService)
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
}
