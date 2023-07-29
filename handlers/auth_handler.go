package handlers

import (
	"os"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/database"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/helpers"
	"github.com/welsonoktario/prakerja-unjuk-ketrampilan/models"
)

func Login(c *fiber.Ctx) error {
	DB := database.Get()

	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	user := models.User{
		Username: payload.Username,
	}

	result := DB.First(&user)

	if result.Error != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + result.Error.Error(),
		})
	}

	if result.RowsAffected == 0 {
		return helpers.ResponseNotFound(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "Username or password is incorrect",
		})
	}

	compareHash, err := argon2id.ComparePasswordAndHash(payload.Password, user.Password)

	if err != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + err.Error(),
		})
	}

	if !compareHash {
		return helpers.ResponseNotFound(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "Username or password is incorrect",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"name":     user.Name,
		"exp":      time.Now().Add(time.Hour * 168).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return helpers.ResponseOK(c, models.JsonResponsePayload{
		Status: "OK",
		Data: fiber.Map{
			"user":  user,
			"token": t,
		},
	})
}

func Register(c *fiber.Ctx) error {
	DB := database.Get()

	payload := struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + err.Error(),
		})
	}

	passwordHash, err := argon2id.CreateHash(payload.Password, argon2id.DefaultParams)

	if err != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + err.Error(),
		})
	}

	user := models.User{
		Username: payload.Username,
		Name:     payload.Name,
		Password: passwordHash,
	}

	result := DB.Create(&user)

	if result.Error != nil {
		return helpers.ResponseFail(c, models.JsonResponsePayload{
			Status:  "FAIL",
			Message: "An error occurred: " + err.Error(),
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"name":     user.Name,
		"exp":      time.Now().Add(time.Hour * 168).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return helpers.ResponseOK(c, models.JsonResponsePayload{
		Message: "OK",
		Data: fiber.Map{
			"user":  user,
			"token": t,
		},
	})
}
