package port

import "github.com/livingdolls/auth-service/internal/core/model"

type UserPortService interface {
	GetDetailUserByEmail(email string) (*model.User, error)
}
