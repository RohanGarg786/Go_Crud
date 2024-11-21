package views

import "github.com/gofiber/fiber/v2"

func JSON(c *fiber.Ctx, status int, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": "success",
		"data":    data,
	})
}
