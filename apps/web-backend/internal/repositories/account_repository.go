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

// GetAccountById retrieves account data by id with last 50 transactions
func (r *AccountRepository) GetAccountById(accountID string, userID string) (*models.AccountDetails, error) {
	query := `
		SELECT 
			a.id, a.account_name, a.user_id, a.total_balance, a.created_at, a.updated_at,
			t.id, t.account_id, t.amount, t.transaction_type, t.transaction_date, t.description, t.created_at, t.updated_at
		FROM accounts a
		LEFT JOIN (
			SELECT id, account_id, amount, transaction_type, transaction_date, description, created_at, updated_at
			FROM transactions
			WHERE account_id = $2
			ORDER BY transaction_date DESC
			LIMIT 50
		) t ON a.id = t.account_id
		WHERE a.user_id = $1 AND a.id = $2`

	rows, err := r.db.Query(query, userID, accountID)
	if err != nil {
		return nil, fmt.Errorf("error getting account with transactions: %v", err)
	}
	defer rows.Close()

	var account *models.Account
	transactions := []models.Transaction{}

	for rows.Next() {
		// If this is the first row, initialize the account
		if account == nil {
			account = &models.Account{}
			var tID, tAccountID, tAmount, tTransactionType, tDescription sql.NullString
			var tDate, tCreatedAt, tUpdatedAt sql.NullTime

			err := rows.Scan(
				&account.ID, &account.AccountName, &account.UserID, &account.TotalBalance, &account.CreatedAt, &account.UpdatedAt,
				&tID, &tAccountID, &tAmount, &tTransactionType, &tDate, &tDescription, &tCreatedAt, &tUpdatedAt)

			if err != nil {
				return nil, fmt.Errorf("error scanning account: %v", err)
			}

			// If there's a transaction, add it
			if tID.Valid {
				transaction := models.Transaction{
					ID:              tID.String,
					AccountID:       tAccountID.String,
					Amount:          tAmount.String,
					TransactionType: tTransactionType.String,
					Date:            tDate.Time,
					Description:     tDescription.String,
					CreatedAt:       tCreatedAt.Time,
					UpdatedAt:       tUpdatedAt.Time,
				}
				transactions = append(transactions, transaction)
			}
		} else {
			// For subsequent rows, only scan transaction data
			var tID, tAccountID, tAmount, tTransactionType, tDescription sql.NullString
			var tDate, tCreatedAt, tUpdatedAt sql.NullTime

			err := rows.Scan(
				&account.ID, &account.AccountName, &account.UserID, &account.TotalBalance, &account.CreatedAt, &account.UpdatedAt,
				&tID, &tAccountID, &tAmount, &tTransactionType, &tDate, &tDescription, &tCreatedAt, &tUpdatedAt)

			if err != nil {
				return nil, fmt.Errorf("error scanning transaction: %v", err)
			}

			if tID.Valid {
				transaction := models.Transaction{
					ID:              tID.String,
					AccountID:       tAccountID.String,
					Amount:          tAmount.String,
					TransactionType: tTransactionType.String,
					Date:            tDate.Time,
					Description:     tDescription.String,
					CreatedAt:       tCreatedAt.Time,
					UpdatedAt:       tUpdatedAt.Time,
				}
				transactions = append(transactions, transaction)
			}
		}
	}

	if account == nil {
		return nil, nil // Account not found
	}

	// Create AccountDetails with embedded Account and transactions
	accountDetails := &models.AccountDetails{
		Account:      *account,
		Transactions: transactions,
	}

	return accountDetails, nil
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

// func (r *AccountRepository) GetAccountDetails(accountID string, email string, provider string, providerID string) (*models.AccountDetails, error) {

// }
