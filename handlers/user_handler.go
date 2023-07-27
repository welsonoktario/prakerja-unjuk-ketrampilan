package handlers

import "github.com/gofiber/fiber/v2"

func GetAllUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}

func GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "OK",
	})
}
