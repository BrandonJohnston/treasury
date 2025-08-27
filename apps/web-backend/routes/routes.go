package routes

import (
	"github.com/gofiber/fiber/v2"

	"web-backend/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/auth")

	auth.Get("/user", handlers.GetUser)
	auth.Post("/login", handlers.LoginUser)
	auth.Post("/register", handlers.CreateUser)
}
