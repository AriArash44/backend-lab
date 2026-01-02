package service

import (
	"errors"

	"1.simple-CRUD/internal/auth/dto"
	"1.simple-CRUD/internal/auth/model"
	"1.simple-CRUD/internal/auth/repository"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(input dto.LoginInput) (*model.User, error) {
	user, err := s.repo.GetByUsername(input.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Password != input.Password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
