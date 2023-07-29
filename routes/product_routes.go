package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/handlers"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/middlewares"
)

func ProductRoutes(app *fiber.App) fiber.Router {
	productRoutes := app.Group("products", middlewares.JwtMiddleware())

	productRoutes.Get("/", handlers.FindAllProducts)
	productRoutes.Get("/:productId", handlers.FindProduct)
	productRoutes.Post("/", handlers.CreateProduct)
	productRoutes.Put("/:productId", handlers.UpdateProduct)
	productRoutes.Delete("/:productId", handlers.DeleteProduct)

	return productRoutes
}
