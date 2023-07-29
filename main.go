package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/database"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/routes"
)

func main() {
	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New()

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Register routes
	routes.Register(app)

	// Listen on port by env APP_PORT value
	log.Fatalln(app.Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))))
}
