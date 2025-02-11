package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func CORS() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientURL := os.Getenv("CLIENT_URL")

		c.Set("Access-Control-Allow-Origin", clientURL)
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Set("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	}
}
