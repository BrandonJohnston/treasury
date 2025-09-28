// internal/handlers/user.go

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"web-backend/internal/models"
	"web-backend/internal/services"
)

type UserHandler struct {
	service *services.UserService
}

type UserData struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Provider   string `json:"provider"`
	ProviderID string `json:"provider_id"`
}

type UserDataResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	User    *models.User `json:"user"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// GetUserData - handle GET requests for user data
func (h *UserHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// For now, return a simple response (you can implement actual user data retrieval)
	response := map[string]interface{}{
		"message": "User data retrieved successfully",
		"status":  "ok",
	}

	json.NewEncoder(w).Encode(response)
}

// PostUserData - handle user data from frontend session
func (h *UserHandler) PostUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Limit the size of the request body to prevent abuse
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // Limit to 1MB

	// Initialize a variable to hold the decoded data
	var userData UserData

	// Decode the request body into the 'data' variable
	// We use a json.Decoder which reads directly from r.Body
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Optional: Returns an error if extra fields are present

	err := decoder.Decode(&userData)
	if err != nil {
		// Handle various decoding errors (e.g., bad JSON, size limit exceeded)
		var msg string
		if err == io.EOF {
			msg = "Request body must not be empty"
		} else if err.Error() == "http: request body too large" {
			msg = "Request body must not be larger than 1MB"
		} else {
			msg = fmt.Sprintf("Bad Request: %v", err)
		}
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Check for extra data after the JSON object (optional but good practice)
	if decoder.More() {
		http.Error(w, "Bad Request: request body contains multiple JSON objects", http.StatusBadRequest)
		return
	}

	// Create a response to use later
	var response UserDataResponse

	// Check if the user already exists
	user, err := h.service.GetUserByEmail(userData.Email, userData.ProviderID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if user != nil {
		response.Status = "ok"
		response.Message = "User already exists"
		response.User = user
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create a new user
	newUser, err := h.service.CreateUser(userData.Provider, userData.ProviderID, userData.Email, userData.Name)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// For now, return the user data back (you can process it with your service later)
	if newUser != nil {
		response.Status = "ok"
		response.Message = "User data received successfully"
		response.User = newUser
	}

	json.NewEncoder(w).Encode(response)
}
