package services

import (
	"web-backend/internal/models"
	"web-backend/internal/repositories"
)

type AccountService struct {
	repo *repositories.AccountRepository
}

func NewAccountService(repo *repositories.AccountRepository) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) GetAccount(accountName string, userID string) (*models.Account, error) {
	return s.repo.GetAccount(accountName, userID)
}

func (s *AccountService) GetAccounts(userID string) ([]models.Account, error) {
	accounts, err := s.repo.GetAccounts(userID)
	if err != nil {
		return nil, err
	}

	// Convert []*models.Account to []models.AccountInfo
	accountInfos := make([]models.Account, len(accounts))
	for i, account := range accounts {
		accountInfos[i] = models.Account{
			ID:           account.ID,
			AccountName:  account.AccountName,
			UserID:       account.UserID,
			TotalBalance: account.TotalBalance,
			CreatedAt:    account.CreatedAt,
			UpdatedAt:    account.UpdatedAt,
		}
	}

	return accountInfos, nil
}

// GetAccountByID retrieves an account by id
func (s *AccountService) GetAccountByID(accountID string, userID string) (*models.AccountDetails, error) {
	return s.repo.GetAccountById(accountID, userID)
}

func (s *AccountService) CreateAccount(accountName string, userID string) (*models.Account, error) {
	return s.repo.CreateAccount(accountName, userID)
}
