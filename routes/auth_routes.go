package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/handlers"
)

func AuthRoutes(app *fiber.App) fiber.Router {
	authRoutes := app.Group("auth")

	authRoutes.Post("/login", handlers.Login)
	authRoutes.Post("/register", handlers.Register)

	return authRoutes
}
