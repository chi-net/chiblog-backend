package routes

import (
	"github.com/gofiber/fiber/v2"
)

func CreateComments(c *fiber.Ctx) error {
	return c.SendString("Create comments")
}