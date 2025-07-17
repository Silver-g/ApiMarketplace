package userservice

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/domain"
	"ApiMarketplace/internal/security"
	"ApiMarketplace/internal/store/repository"
	"context"
)

type UserRegister interface {
	Register(ctx context.Context, user domain.RegisterUserInternal) (*boundary.RegisterUserResponse, error)
}
type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}
func (s *UserService) Register(ctx context.Context, user domain.RegisterUserInternal) (*boundary.RegisterUserResponse, error) {
	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	CreateUserData := boundary.RegisterUserDbMaping(user)
	createdUser, err := s.userRepo.CreateUser(ctx, &CreateUserData)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
