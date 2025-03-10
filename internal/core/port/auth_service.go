package port

import "github.com/livingdolls/auth-service/internal/core/model"

type AuthPortService interface {
	Register(username, email, passwird string) (*model.User, error)
	Login(email, password string) (string, error)
	Logout(UserID string) error
}
