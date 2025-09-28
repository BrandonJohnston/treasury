package repositories

import (
	"database/sql"
	"fmt"

	"web-backend/internal/models"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUserByEmail retrieves a user by email & provider_id
func (r *UserRepository) GetUserByEmail(email string, providerID string) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT id, provider, provider_id, email, name, created_at, updated_at
		FROM users
		WHERE email = $1
		AND provider_id = $2`

	err := r.db.QueryRow(query, email, providerID).Scan(
		&user.ID, &user.Provider, &user.ProviderID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("error getting user by email: %v", err)
	}

	return user, nil
}

// GetUserByID retrieves a user by id
func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}

	query := `
		SELECT id, provider, provider_id, email, name, created_at, updated_at
		FROM users
		WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Provider, &user.ProviderID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("error getting user by id: %v", err)
	}

	return user, nil
}

// CreateUser creates a new user
func (r *UserRepository) CreateUser(provider string, providerID string, email string, name string) (*models.User, error) {
	user := &models.User{}

	query := `
		INSERT INTO users (provider, provider_id, email, name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, provider, provider_id, email, name, created_at, updated_at`

	err := r.db.QueryRow(query, provider, providerID, email, name).Scan(
		&user.ID, &user.Provider, &user.ProviderID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return user, nil
}
