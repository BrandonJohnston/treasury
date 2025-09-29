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

func (s *UserService) CreateAccount(provider string) (*models.Account, error) {
	return s.repo.CreateAccount(provider)
}
