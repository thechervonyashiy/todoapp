package services

import (
	"errors"

	"github.com/thechervonyashiy/todoapp/pkg/dto"
	"github.com/thechervonyashiy/todoapp/pkg/models"
	"github.com/thechervonyashiy/todoapp/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(req dto.SignUpRequest) (*dto.SignUpResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	response, err := s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *AuthService) SignIn(req dto.SignInRequest) (*dto.SignInResponse, error) {
	return nil, nil
}
