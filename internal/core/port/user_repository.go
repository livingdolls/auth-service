package port

import "github.com/livingdolls/auth-service/internal/core/model"

type UserRepository interface {
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) error
}
