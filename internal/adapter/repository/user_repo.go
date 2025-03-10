package repository

import (
	"database/sql"

	"github.com/livingdolls/auth-service/internal/core/model"
	"github.com/livingdolls/auth-service/internal/core/port"
)

type userRepository struct {
	db *sql.DB
}

// Create implements port.UserRepository.
func (r *userRepository) Create(user *model.User) error {
	query := `INSERT INTO users (id, username, email, password_hash, created_at) VALUES ($1, $2, $3, $4, NOW())`
	_, err := r.db.Exec(query, user.ID, user.Username, user.Email, user.PasswordHash)
	return err
}

// GetByEmail implements port.UserRepository.
func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT id, username, email, password_hash FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *sql.DB) port.UserRepository {
	return &userRepository{db: db}
}
