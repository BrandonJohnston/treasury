package models

import "time"

// User represents a user in the database
type Account struct {
	ID         string    `json:"id" db:"id"`
	Provider   string    `json:"provider" db:"provider"`
	ProviderID string    `json:"provider_id" db:"provider_id"`
	Email      string    `json:"email" db:"email"`
	Name       string    `json:"name" db:"name"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// CreateUserRequest represents the structure for creating a new user
type CreateAccountRequest struct {
	AccountName string `json:"account_name" db:"account_name"`
}
