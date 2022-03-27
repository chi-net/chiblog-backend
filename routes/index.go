package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Index (c *fiber.Ctx) {
	return c.SendString("Hello, World 👋!")
}