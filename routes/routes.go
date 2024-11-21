package routes

import (
	"Crud_fiber_Go/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User Routes
	api.Get("/users", controllers.GetUsers)
	api.Post("/users", controllers.CreateUser)
}
