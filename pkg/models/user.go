package models

import "time"

type User struct {
	Id        int       `json:"-"`
	Username  string    `json:"username" binding:"required"` //binding:"required" теги валидируют наличие этих полей в запросе
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
