package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"web-backend/internal/database"
	"web-backend/internal/handlers"
	"web-backend/internal/repositories"
	"web-backend/internal/router"
	"web-backend/internal/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	db, err := database.SetupDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	routes := router.SetupRoutes(userHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
