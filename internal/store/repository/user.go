package repository

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/domain"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.RegisterUserDB) (*boundary.RegisterUserResponse, error)
}
