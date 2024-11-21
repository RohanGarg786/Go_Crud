package main

import (
	"Crud_fiber_Go/config"
	"Crud_fiber_Go/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration using Builder Pattern
	appConfig := config.NewAppConfigBuilder().
		SetPort(":3000").
		SetDatabase("postgres://password@localhost:5432/Go_database").
		Build()

	// Initialize Fiber App
	app := fiber.New()

	// Connect to database
	if err := appConfig.ConnectDB(); err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Register routes
	routes.RegisterRoutes(app)

	// Start server
	log.Fatal(app.Listen(appConfig.Port))
}
