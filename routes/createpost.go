package routes

import (
	"github.com/gofiber/fiber/v2"
)

func CreatePosts(c *fiber.Ctx) error {
	return c.SendString("Posts New")
}