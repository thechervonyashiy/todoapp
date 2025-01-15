package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/thechervonyashiy/todoapp/pkg/dto"
	"github.com/thechervonyashiy/todoapp/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (*dto.SignUpResponse, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
