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
	service     *services.AccountService
	userService *services.UserService
}

type NewAccountData struct {
	AccountName string `json:"accountName"`
	Email       string `json:"email"`
	Provider    string `json:"provider"`
	ProviderID  string `json:"providerId"`
}

type AccountDataResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Account *models.Account `json:"account"`
}

type GetAccountsResponse struct {
	Status   string               `json:"status"`
	Message  string               `json:"message"`
	Accounts []models.AccountInfo `json:"accounts"`
}

func NewAccountHandler(service *services.AccountService, userService *services.UserService) *AccountHandler {
	return &AccountHandler{
		service:     service,
		userService: userService,
	}
}

// GetAccounts - handle GET requests for all accounts data
func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract query parameters
	email := r.URL.Query().Get("email")
	provider := r.URL.Query().Get("provider")
	providerID := r.URL.Query().Get("provider_id")

	// Validate required parameters
	if email == "" || provider == "" || providerID == "" {
		response := GetAccountsResponse{
			Status:   "error",
			Message:  "Missing required query parameters: email, provider, provider_id",
			Accounts: []models.AccountInfo{},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get the user ID first
	user, err := h.userService.GetUserByEmail(email, providerID)
	if err != nil {
		response := GetAccountsResponse{
			Status:   "error",
			Message:  "Internal Server Error",
			Accounts: []models.AccountInfo{},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if user == nil {
		response := GetAccountsResponse{
			Status:   "error",
			Message:  "User not found",
			Accounts: []models.AccountInfo{},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get the accounts for the user
	accounts, err := h.service.GetAccounts(user.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// For now, return a simple response (you can implement actual account data retrieval)
	response := GetAccountsResponse{
		Message:  "Account data retrieved successfully",
		Status:   "ok",
		Accounts: accounts,
	}

	json.NewEncoder(w).Encode(response)
}

// PostAccountData - handle new account data from frontend session
func (h *AccountHandler) PostAccountData(w http.ResponseWriter, r *http.Request) {
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

	// Get the user ID first
	user, err := h.userService.GetUserByEmail(accountData.Email, accountData.ProviderID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check if the Account already exists
	account, err := h.service.GetAccount(accountData.AccountName, user.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if account != nil {
		response.Status = "err"
		response.Message = "Account already exists"
		response.Account = account
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create a new Account
	newAccount, err := h.service.CreateAccount(accountData.AccountName, user.ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// For now, return the user data back (you can process it with your service later)
	if newAccount != nil {
		response.Status = "ok"
		response.Message = "User account created successfully"
		response.Account = newAccount
	}

	json.NewEncoder(w).Encode(response)
}
