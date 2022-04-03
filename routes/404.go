package routes

import (
	"github.com/gofiber/fiber/v2"
)

func F0f(c *fiber.Ctx) error {
	return c.Status(401).SendString("")
}
