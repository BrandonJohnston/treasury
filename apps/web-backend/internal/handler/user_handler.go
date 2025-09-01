// internal/handler/user_handler.go

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"

	"web-backend/internal/service"
)

type UserHandler struct {
	service      *service.UserService
	sessionStore *sessions.CookieStore
}

type AuthUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewUserHandler(service *service.UserService, sessionStore *sessions.CookieStore) *UserHandler {
	return &UserHandler{
		service:      service,
		sessionStore: sessionStore,
	}
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
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

	session, err := h.sessionStore.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	// Set session values
	session.Values["authenticated"] = true
	session.Values["userID"] = user.ID
	session.Save(r, w)

	json.NewEncoder(w).Encode(user)
}

// GetUserData - return user data from active session
func (h *UserHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	session, err := h.sessionStore.Get(r, "session-name")

	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	// Check if user is authenticated
	auth, ok := session.Values["authenticated"].(bool)

	// User is not authenticated
	if !ok || !auth {
		// Write the error response to the client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		// Send JSON response
		response := ErrorResponse{
			Message: "Unathenticated",
			Status:  http.StatusUnauthorized,
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// Get user ID from session
	userID, ok := session.Values["userID"].(int)
	if !ok {
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

	// Retrieve user from the service layer
	user, err := h.service.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
