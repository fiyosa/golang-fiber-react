package config

import (
	"github.com/gofiber/fiber/v2"
)

func Cors() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set(
			"Access-Control-Allow-Headers", 
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, origin, Cache-Control, X-Requested-With",
		)
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Method() == "OPTIONS" {
			return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Method is not permitted"})
		}

		return c.Next()
	}
}
