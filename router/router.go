package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"hanhngo.me/m/modules/users"
)

func SetupRouters(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	api := app.Group("/api", logger.New())

	userRoutes := api.Group("/users", logger.New())
	userRoutes.Get("/", users.GetUserListHandler)
	userRoutes.Get("/:id", users.GetUserByIdHandler)
	userRoutes.Patch("/:id", users.UpdateUserProfileHandler)
	userRoutes.Delete("/:id", users.DeleteUserHandler)
}
