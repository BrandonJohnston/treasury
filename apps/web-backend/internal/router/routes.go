package router

import (
	"net/http"
	"web-backend/internal/handler"
	"web-backend/internal/middleware"
)

func SetupRoutes(userHandler *handler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/auth/login", userHandler.LoginUser)
	mux.HandleFunc("/api/auth/userdata", userHandler.GetUserData)

	return middleware.CORS(mux)
}
