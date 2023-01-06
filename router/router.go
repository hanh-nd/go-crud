package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouters(app *fiber.App) {
	ping := app.Group("/", logger.New())

	ping.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
