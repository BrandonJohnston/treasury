package router

import (
	"net/http"
	"web-backend/internal/handlers"
	"web-backend/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(userHandler *handlers.UserHandler, accountHandler *handlers.AccountHandler) http.Handler {
	r := mux.NewRouter()

	// Define routes with specific HTTP methods
	r.HandleFunc("/api/user", userHandler.GetUserData).Methods("GET")
	r.HandleFunc("/api/user", userHandler.PostUserData).Methods("POST")

	// Account routes
	r.HandleFunc("/api/accounts/create", accountHandler.PostAccountData).Methods("POST")

	return middleware.CORS(r)
}
