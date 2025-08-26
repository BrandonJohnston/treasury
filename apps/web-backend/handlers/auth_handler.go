package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"web-backend/models"
)

func GetUser(c *fiber.Ctx) error {
	log.Println("GET /api/auth/user")

	users := []models.User{
		{ID: 1, Name: "Corben Dallas"},
	}
	return c.JSON(users)
}

func LoginUser(c *fiber.Ctx) error {

	log.Println("POST /api/auth/login")

	var loginReq models.LoginRequest

	// Parse the request body
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	log.Println("Email: ", loginReq.Email)
	log.Println("Password: ", loginReq.Password)

	// Return the email and password values
	return c.JSON(fiber.Map{
		"email":    loginReq.Email,
		"password": loginReq.Password,
	})
}
