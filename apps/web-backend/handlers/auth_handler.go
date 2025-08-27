package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"web-backend/models"
	"web-backend/repository"
)

// UserRepository instance
var userRepo = repository.NewUserRepository()

func GetUser(c *fiber.Ctx) error {
	log.Println("GET /api/auth/user")

	return c.JSON(fiber.Map{
		"users": []fiber.Map{
			{"id": 1, "name": "John Doe", "email": "john@example.com"},
		},
	})
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

	// Get user from database
	user, err := userRepo.GetUserByEmail(loginReq.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	if user == nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// TODO: Add password hashing and verification
	// For now, just return the user info
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		},
	})
}

func CreateUser(c *fiber.Ctx) error {
	log.Println("POST /api/auth/register")

	var createReq models.CreateUserRequest

	// Parse the request body
	if err := c.BodyParser(&createReq); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if user already exists
	existingUser, err := userRepo.GetUserByEmail(createReq.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	if existingUser != nil {
		return c.Status(409).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	// TODO: Add password hashing
	// For now, store password as-is (NOT recommended for production)
	user := &models.User{
		Email:        createReq.Email,
		PasswordHash: createReq.Password, // This should be hashed
		Name:         createReq.Name,
	}

	err = userRepo.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		},
	})
}
