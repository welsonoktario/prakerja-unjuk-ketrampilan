package middlewares

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/helpers"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/models"
)

// Middleware JWT function
func JwtMiddleware() fiber.Handler {
	jwtSecret := os.Getenv("JWT_SECRET")

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return helpers.ResponseUnauthorized(c, models.JsonResponsePayload{
				Status:  "FAIL",
				Message: "Unauthorized",
			})
		},
	})
}
