package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"hanhngo.me/m/config"
	"hanhngo.me/m/database"
	"hanhngo.me/m/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.Connect()

	router.SetupRouters(app)

	port := config.Get("PORT", "3000")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
