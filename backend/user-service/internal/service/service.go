package service

import (
	"context"
	"errors"

	"github.com/abenezer54/gojo/backend/user-service/internal/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repository model.UserRepository
}

func NewUserService(r model.UserRepository) model.UserService {
	return &UserService{
		repository: r,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, request *model.SignupRequest) error {
	_, err := s.repository.GetUserByEmail(ctx, request.Email)

	if err == nil {
		return errors.New("email already in use")
	}

	// If the error is not - "not found" which could be db or something internal error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &model.User{
		ID:         uuid.New(),
		FullName:   request.FullName,
		Email:      request.Email,
		Password:   string(hashedPassword),
		Role:       request.Role,
		IsVerified: false,
	}

	err = s.repository.CreateUser(ctx, newUser)
	return err
}
