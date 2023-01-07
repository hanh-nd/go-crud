package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "hanhngo.me/m/handler"
)

func SetupRouters(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	api := app.Group("/api", logger.New())

	api.Group("/users", logger.New())
}
