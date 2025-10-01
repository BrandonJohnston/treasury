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

func (s *AccountService) GetAccounts(userID string) ([]models.AccountInfo, error) {
	accounts, err := s.repo.GetAccounts(userID)
	if err != nil {
		return nil, err
	}

	// Convert []*models.Account to []models.AccountInfo
	accountInfos := make([]models.AccountInfo, len(accounts))
	for i, account := range accounts {
		accountInfos[i] = models.AccountInfo{
			ID:          account.ID,
			AccountName: account.AccountName,
			UserID:      account.UserID,
			CreatedAt:   account.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   account.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	return accountInfos, nil
}

func (s *AccountService) CreateAccount(accountName string, userID string) (*models.Account, error) {
	return s.repo.CreateAccount(accountName, userID)
}
