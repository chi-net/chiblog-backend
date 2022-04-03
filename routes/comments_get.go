package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetComments(c *fiber.Ctx) error {
	return c.SendString("List comments")
}