package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	AuthRoutes(app)
	ProductRoutes(app)
}
