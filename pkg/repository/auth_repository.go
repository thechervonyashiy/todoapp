package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/thechervonyashiy/todoapp/pkg/dto"
	"github.com/thechervonyashiy/todoapp/pkg/models"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) (*dto.SignUpResponse, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)

	var userID int
	if err := row.Scan(&userID); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	response := &dto.SignUpResponse{
		ID:       userID,
		Username: user.Username,
	}

	return response, nil
}

func (r *AuthRepository) GetUser(username, password string) (models.User, error) {
	return models.User{}, nil
}
