package controllers

import (
	"Crud_fiber_Go/models"
	"Crud_fiber_Go/views"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	// Example: Fetch users from database
	var users []models.User
	// DB logic here...

	return views.JSON(c, fiber.StatusOK, users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return views.JSON(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Save user to database
	// DB logic here...

	return views.JSON(c, fiber.StatusCreated, user)
}
