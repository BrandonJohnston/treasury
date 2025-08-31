package service

import (
	"log"
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
	log.Println("Service :: GetUserByEmail()")
	return s.repo.GetUserByEmail(email, password)
}
