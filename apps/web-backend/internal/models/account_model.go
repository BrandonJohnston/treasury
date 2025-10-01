package models

import "time"

// User represents a user in the database
type Account struct {
	ID          string    `json:"id" db:"id"`
	AccountName string    `json:"accountName" db:"account_name"`
	UserID      string    `json:"userId" db:"user_id"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

// CreateUserRequest represents the structure for creating a new user
type CreateAccountRequest struct {
	AccountName string `json:"accountName" db:"account_name"`
	Email       string `json:"email" db:"email"`
	Provider    string `json:"provider" db:"provider"`
	ProviderID  string `json:"providerId" db:"provider_id"`
}

// AccountInfo represents account information for API responses
type AccountInfo struct {
	ID          string `json:"id"`
	AccountName string `json:"accountName"`
	UserID      string `json:"userId"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
