package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/chihuo2104/chiblog-backend/routes"
)

func main() {
	app := fiber.New()

	app.Get("/", routes.Index())

	app.Listen(":3000")
}
