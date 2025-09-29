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

func (s *AccountService) CreateAccount(accountName string, userID string) (*models.Account, error) {
	return s.repo.CreateAccount(accountName, userID)
}
