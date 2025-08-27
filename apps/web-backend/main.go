package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"web-backend/config"
	"web-backend/database"
	"web-backend/routes"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	app := fiber.New()

	// Enable CORS for all routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	routes.SetupRoutes(app)

	log.Println("Server starting on :8080")
	log.Fatal(app.Listen(":8080"))
}
