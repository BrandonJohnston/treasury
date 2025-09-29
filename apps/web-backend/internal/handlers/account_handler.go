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

type AccountHandler struct {
	service *services.AccountService
}

type NewAccountData struct {
	AccountName string `json:"account_name"`
}

type AccountDataResponse struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	User    *models.User `json:"user"`
}

func NewAccountHandler(service *services.AccountService) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}

// PostAccountData - handle new account data from frontend session
func (h *UserHandler) PostAccountData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Limit the size of the request body to prevent abuse
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // Limit to 1MB

	// Initialize a variable to hold the decoded data
	var accountData NewAccountData

	// Decode the request body into the 'data' variable
	// We use a json.Decoder which reads directly from r.Body
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Optional: Returns an error if extra fields are present

	err := decoder.Decode(&accountData)
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
	var response AccountDataResponse

	json.NewEncoder(w).Encode(response)
}
