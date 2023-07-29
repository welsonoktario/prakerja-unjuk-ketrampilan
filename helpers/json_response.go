package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/models"
)

func ResponseOK(c *fiber.Ctx, data models.JsonResponsePayload) error {
	return c.Status(fiber.StatusOK).JSON(data)
}

func ResponseCreated(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(data)
}

func ResponseNotFound(c *fiber.Ctx, data models.JsonResponsePayload) error {
	return c.Status(fiber.StatusNotFound).JSON(data)
}

func ResponseFail(c *fiber.Ctx, data models.JsonResponsePayload) error {
	return c.Status(fiber.StatusInternalServerError).JSON(data)
}

func ResponseUnauthorized(c *fiber.Ctx, data models.JsonResponsePayload) error {
	return c.Status(fiber.StatusUnauthorized).JSON(data)
}
