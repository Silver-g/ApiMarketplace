package repository

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/domain"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.RegisterUserDb) (*boundary.RegisterUserResponse, error)
	LoginByUsername(ctx context.Context, username *domain.LoginUserDb) (*domain.LoginUserResponseDb, error)
}
