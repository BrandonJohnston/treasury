package service

import (
	"web-backend/internal/models"
	"web-backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByEmail(email string, password string) (*models.User, error) {
	return s.repo.GetUserByEmail(email, password)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
