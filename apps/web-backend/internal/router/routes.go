package router

import (
	"net/http"
	"web-backend/internal/handlers"
	"web-backend/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(userHandler *handlers.UserHandler) http.Handler {
	r := mux.NewRouter()

	// Define routes with specific HTTP methods
	r.HandleFunc("/api/user", userHandler.GetUserData).Methods("GET")
	r.HandleFunc("/api/user", userHandler.PostUserData).Methods("POST")

	return middleware.CORS(r)
}
