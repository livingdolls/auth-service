package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/livingdolls/auth-service/internal/core/model"
	"github.com/livingdolls/auth-service/internal/core/port"
	"github.com/livingdolls/auth-service/internal/infrastructure/security"
)

type AuthService struct {
	userRepo port.UserRepository
}

func NewAuthService(userRepo port.UserRepository) port.AuthPortService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(username, email, password string) (*model.User, error) {
	existingUser, _ := s.userRepo.GetByEmail(email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, _ := security.HashPassword(password)
	user := &model.User{
		ID:           uuid.New().String(),
		Username:     username,
		Email:        email,
		PasswordHash: hashedPassword,
		CreatedAt:    time.Now(),
	}

	err := s.userRepo.Create(user)
	return user, err
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil || !security.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	token, err := security.GenerateJWT(user.ID)
	return token, err
}

func (s *AuthService) Logout(UserID string) error {
	return nil
}
