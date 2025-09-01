package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"

	"web-backend/internal/database"
	"web-backend/internal/handler"
	"web-backend/internal/repository"
	"web-backend/internal/router"
	"web-backend/internal/service"
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

	// Create session store (secret should be from environment in production)
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		sessionKey = "default-insecure-key" // Fallback key (not for production!)
	}
	sessionStore := sessions.NewCookieStore([]byte(sessionKey))
	sessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600, // 1 hour
		HttpOnly: true,
		// Secure: true,	// Enable ONLY when using HTTPS
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, sessionStore)

	routes := router.SetupRoutes(userHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
