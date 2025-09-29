package repositories

import (
	"database/sql"
	// "fmt"

	"web-backend/internal/models"
)

// UserRepository handles database operations for users
type AccountRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// CreateUser creates a new user
func (r *AccountRepository) CreateAccount(provider string) (*models.Account, error) {
	user := &models.Account{}

	// query := `
	// 	INSERT INTO users (provider, provider_id, email, name)
	// 	VALUES ($1, $2, $3, $4)
	// 	RETURNING id, provider, provider_id, email, name, created_at, updated_at`

	// err := r.db.QueryRow(query, provider, providerID, email, name).Scan(
	// 	&user.ID, &user.Provider, &user.ProviderID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	// if err != nil {
	// 	return nil, fmt.Errorf("error creating user: %v", err)
	// }

	return user, nil
}
