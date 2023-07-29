package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/database"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/helpers"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/models"
)

func FindAllProducts(c *fiber.Ctx) error {
	DB := database.Get()
	products := []models.Product{}

	own, err := strconv.ParseBool(c.Query("own", "false"))

	if err != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + err.Error(),
		})
	}

	result := DB.Model(&products).Preload("User").Find(&products)
	if own {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := uint(claims["id"].(float64))

		result = result.Where("user_id = ?", userId)
	}

	if result.Error != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + result.Error.Error(),
		})
	}

	return helpers.ResponseOK(c, models.JsonResponsePayload{
		Status: "OK",
		Data: fiber.Map{
			"products": products,
		},
	})
}

func FindProduct(c *fiber.Ctx) error {
	DB := database.Get()

	productId, err := strconv.Atoi(c.Params("productId"))

	if err != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "Provided 'productId' parameter must be numeric",
		})
	}

	product := models.Product{
		ID: uint(productId),
	}

	result := DB.Model(&product).Preload("User").First(&product)

	if result.Error != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + err.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return helpers.ResponseNotFound(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: fmt.Sprintf("Product with id '%d' not found", productId),
		})
	}

	return helpers.ResponseOK(c, models.JsonResponsePayload{
		Status: "OK",
		Data: fiber.Map{
			"product": product,
		},
	})
}

func CreateProduct(c *fiber.Ctx) error {
	DB := database.Get()
	authUser := c.Locals("user").(*jwt.Token)
	claims := authUser.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	payload := struct {
		Name        string `json:"name"`
		Price       uint   `json:"price"`
		Description string `json:"description"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	product := models.Product{
		Name:        payload.Name,
		Price:       payload.Price,
		Description: &payload.Description,
		UserID:      uint(userId),
	}

	result := DB.Model(&product).Preload("User").Create(&product)

	if result.Error != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + result.Error.Error(),
		})
	}

	return helpers.ResponseOK(c, models.JsonResponsePayload{
		Status:  "OK",
		Message: "Product successfully created",
		Data: fiber.Map{
			"product": product,
		},
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	DB := database.Get()
	authUser := c.Locals("user").(*jwt.Token)
	claims := authUser.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	productId, err := strconv.Atoi(c.Params("productId"))

	if err != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "Provided 'productId' parameter must be numeric",
		})
	}

	payload := struct {
		Name        string `json:"name"`
		Price       uint   `json:"price"`
		Description string `json:"description"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	product := models.Product{
		ID:          uint(productId),
		Name:        payload.Name,
		Price:       payload.Price,
		Description: &payload.Description,
		UserID:      userId,
	}

	result := DB.Model(&product).Preload("User").Updates(&product)

	if result.Error != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + result.Error.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return helpers.ResponseNotFound(c, models.JsonResponsePayload{
			Status:  "OK",
			Message: fmt.Sprintf("Product with id '%d' not found", productId),
		})
	}

	return helpers.ResponseOK(c, models.JsonResponsePayload{
		Status:  "OK",
		Message: "Product successfully updated",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	DB := database.Get()
	authUser := c.Locals("user").(*jwt.Token)
	claims := authUser.Claims.(jwt.MapClaims)
	userId := claims["id"].(float64)

	productId, err := strconv.Atoi(c.Params("productId"))

	if err != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "Provided 'productId' parameter must be numeric",
		})
	}

	product := models.Product{
		ID:     uint(productId),
		UserID: uint(userId),
	}

	result := DB.Model(&product).Delete(&product)

	if result.Error != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + err.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return helpers.ResponseNotFound(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: fmt.Sprintf("Product with id '%d' not found", productId),
		})
	}

	return helpers.ResponseOK(c, models.JsonResponsePayload{
		Status:  "OK",
		Message: "Product successfully deleted",
	})
}
