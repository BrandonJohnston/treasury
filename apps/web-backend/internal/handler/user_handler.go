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
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
