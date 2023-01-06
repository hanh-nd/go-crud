package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"hanhngo.me/m/handler"
)

func SetupRouters(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	api := app.Group("/api", logger.New())

	user := api.Group("/users", logger.New())
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)
}
