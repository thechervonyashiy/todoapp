package services

import (
	"github.com/thechervonyashiy/todoapp/pkg/dto"
	"github.com/thechervonyashiy/todoapp/pkg/repository"
)

type Authorization interface {
	CreateUser(req dto.SignUpRequest) (*dto.SignUpResponse, error)
	SignIn(req dto.SignInRequest) (*dto.SignInResponse, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}
