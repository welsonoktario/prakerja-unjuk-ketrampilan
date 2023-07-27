package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/database"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/handlers"
)

func main() {
	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/users", handlers.GetAllUsers)
	v1.Post("/users", handlers.GetUser)

	// Listen on port 3000
	log.Fatal(app.Listen(":8080")) // go run app.go -port=:3000
}
