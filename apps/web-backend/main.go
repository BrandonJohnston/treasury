package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"web-backend/routes"
)

func main() {
	app := fiber.New()

	// Enable CORS for all routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
