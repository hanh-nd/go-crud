package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"hanhngo.me/m/modules/auth"
	"hanhngo.me/m/modules/users"
)

var (
	authService auth.AuthService  = auth.AuthService{}
	authHandler auth.AuthHandler  = auth.NewAuthHandler(authService, userService)
	userService users.UserService = users.UserService{}
	userHandler users.UserHandler = users.NewUserHandler(userService)
)

func SetupRouters(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	api := app.Group("/api", logger.New())

	authRoutes := app.Group("/", logger.New())
	authRoutes.Post("/register")
	authRoutes.Post("/login")

	userRoutes := api.Group("/users", logger.New())
	userRoutes.Get("/", userHandler.GetUserList)
	userRoutes.Get("/:id", userHandler.GetUserById)
	userRoutes.Patch("/:id", userHandler.UpdateUserProfile)
	userRoutes.Delete("/:id", userHandler.DeleteUser)
}
