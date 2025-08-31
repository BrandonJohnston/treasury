package repository

import (
	"database/sql"
	"fmt"
	"log"

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

// GetUserByEmail retrieves a user by email
func (r *UserRepository) GetUserByEmail(email string, password string) (*models.User, error) {
	log.Println("Repository :: GetUserByEmail()")

	log.Println("email: ", email)
	log.Println("password: ", password)

	user := &models.User{}
	query := `
		SELECT id, email, password_hash, name, created_at, updated_at
		FROM users
		WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, fmt.Errorf("error getting user by email: %v", err)
	}

	return user, nil
}
