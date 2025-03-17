package service

import (
	"github.com/livingdolls/auth-service/internal/core/model"
	"github.com/livingdolls/auth-service/internal/core/port"
)

type UserService struct {
	userRepo port.UserRepository
}

// GetDetailUserByEmail implements port.UserPortService.
func (u *UserService) GetDetailUserByEmail(email string) (*model.User, error) {
	user, err := u.userRepo.GetByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserService(userRepo port.UserRepository) port.UserPortService {
	return &UserService{userRepo: userRepo}
}
