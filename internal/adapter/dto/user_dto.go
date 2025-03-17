package dto

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserGetByEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}
