package handler

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	c.JSON(fiber.Map{
		"status": "ok",
	})
	return nil
}
