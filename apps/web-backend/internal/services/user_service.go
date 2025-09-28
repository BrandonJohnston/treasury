package services

import (
	"web-backend/internal/models"
	"web-backend/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByEmail(email string, providerID string) (*models.User, error) {
	return s.repo.GetUserByEmail(email, providerID)
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(provider string, providerID string, email string, name string) (*models.User, error) {
	return s.repo.CreateUser(provider, providerID, email, name)
}
