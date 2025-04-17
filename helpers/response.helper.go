package helpers

import (
	"github.com/gofiber/fiber/v2"
)
func RespondWithError(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"success": false,
		"error":   message,
	})
}

func RespondWithSuccess(c *fiber.Ctx, status int, message interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"success": true,
		"data":    message,
	})
}