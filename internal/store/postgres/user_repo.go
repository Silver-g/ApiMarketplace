package postgres

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/consts"
	"ApiMarketplace/internal/domain"
	"context"
	"database/sql"
	"errors"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(dataBase *sql.DB) *UserPostgres {
	return &UserPostgres{db: dataBase}
}

var (
	ErrUserAlreadyExists = errors.New(consts.ErrUserAlreadyExistsMsg)
	ErrUserNotFound      = errors.New(consts.ErrUserNotFoundMsg)
)

func (r *UserPostgres) LoginByUsername(ctx context.Context, username *domain.LoginUserDB) (*domain.LoginUserResponseDb, error) {
	var dbLoginResponse domain.LoginUserResponseDb
	query := "SELECT id, password_hash FROM users WHERE login = $1"
	err := r.db.QueryRowContext(ctx, query, username.Username).Scan(&dbLoginResponse.Id, &dbLoginResponse.PasswordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &dbLoginResponse, nil
}

func (r *UserPostgres) CreateUser(ctx context.Context, userData *domain.RegisterUserDB) (*boundary.RegisterUserResponse, error) {

	query := "INSERT INTO users (login, password_hash) VALUES ($1, $2) ON CONFLICT (login) DO NOTHING RETURNING id"
	err := r.db.QueryRowContext(ctx, query, userData.Username, userData.PasswordHash).Scan(&userData.Id)
	if userData.Id == 0 {
		return nil, ErrUserAlreadyExists
	}
	if err != nil {
		return nil, err
	}
	respData := boundary.RegisterUserResponseMaping(*userData)
	return &respData, nil
}
