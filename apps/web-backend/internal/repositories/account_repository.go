package repositories

import (
	"database/sql"
	"fmt"

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

// GetAccount retrieves an account by email, provider, & provider_id
func (r *AccountRepository) GetAccount(accountName string, userID string) (*models.Account, error) {
	account := &models.Account{}

	query := `
		SELECT id, account_name, user_id, total_balance, created_at, updated_at
		FROM accounts
		WHERE account_name = $1
		AND user_id = $2`

	err := r.db.QueryRow(query, accountName, userID).Scan(
		&account.ID, &account.AccountName, &account.UserID, &account.TotalBalance, &account.CreatedAt, &account.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Account not found
		}
		return nil, fmt.Errorf("error getting account by name and user_id: %v", err)
	}

	return account, nil
}

// GetAccounts retrieves all accounts for a user
func (r *AccountRepository) GetAccounts(userID string) ([]*models.Account, error) {
	accounts := []*models.Account{}

	query := `
		SELECT id, account_name, user_id, total_balance, created_at, updated_at
		FROM accounts
		WHERE user_id = $1`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting accounts: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		account := &models.Account{}
		err := rows.Scan(&account.ID, &account.AccountName, &account.UserID, &account.TotalBalance, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning account: %v", err)
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// CreateUser creates a new user
func (r *AccountRepository) CreateAccount(accountName string, userID string) (*models.Account, error) {
	user := &models.Account{}

	query := `
		INSERT INTO accounts (account_name, user_id)
		VALUES ($1, $2)
		RETURNING id, account_name, user_id, created_at, updated_at`

	err := r.db.QueryRow(query, accountName, userID).Scan(
		&user.ID, &user.AccountName, &user.UserID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return user, nil
}
