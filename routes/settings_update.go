package routes

import (
	"github.com/gofiber/fiber/v2"
)

func UpdateSettings(c *fiber.Ctx) error {
	return c.SendString("update settings")
}