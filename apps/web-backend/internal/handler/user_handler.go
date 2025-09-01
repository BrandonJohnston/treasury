// internal/handler/user_handler.go

package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"web-backend/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

type AuthUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("LoginUser called with method: %s", r.Method)
	log.Println("Handler :: LoginUser()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req AuthUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByEmail(req.Email, req.Password)

	// Error on db query
	if err != nil {
		// Write the error response to the client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		// Send JSON response
		response := ErrorResponse{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// User not found
	if user == nil {
		// Write the error response to the client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		// Send JSON response
		response := ErrorResponse{
			Message: "Invalid username or password",
			Status:  http.StatusUnauthorized,
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(user)
}
