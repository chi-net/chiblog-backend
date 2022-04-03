package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetSettings(c *fiber.Ctx) error {
	return c.SendString("Get Settings")
}