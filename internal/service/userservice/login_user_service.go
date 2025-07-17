package userservice

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/domain"
	"ApiMarketplace/internal/security"
	"ApiMarketplace/internal/store/postgres"
	"context"
)

type UserLogin interface {
	UserLogin(ctx context.Context, loginData domain.LoginUserInternal) (boundary.LoginUserResponse, error)
}

func (s *UserService) UserLogin(ctx context.Context, loginData domain.LoginUserInternal) (boundary.LoginUserResponse, error) {
	var err error
	var loginToken boundary.LoginUserResponse
	dbLoginData := boundary.LoginUserDbMaping(loginData)
	userData, err := s.userRepo.LoginByUsername(ctx, &dbLoginData)
	if err != nil {
		if err == postgres.ErrUserNotFound {
			return boundary.LoginUserResponse{}, err
		}
		return boundary.LoginUserResponse{}, err
	}
	err = security.ComparePassword(userData.PasswordHash, loginData.Password)
	if err != nil {
		return boundary.LoginUserResponse{}, err
	}
	token, err := security.GenerateJWT(userData.Id)
	if err != nil {
		return boundary.LoginUserResponse{}, err
	}
	loginToken = boundary.LoginUserResponseMapping(token)
	return loginToken, nil
}
