package userservice

import (
	"ApiMarketplace/internal/domain"
	"ApiMarketplace/internal/mocks"
	"ApiMarketplace/internal/security"
	"ApiMarketplace/internal/store/postgres"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

type loginTestCase struct {
	name    string
	input   domain.LoginUserInternal
	setup   func(m *mocks.MockUserRepository)
	wantErr error
}

func TestLoginUser(t *testing.T) {
	cases := []loginTestCase{
		{
			name: "ValidLoginTest",
			input: domain.LoginUserInternal{
				Username: "ValidUsername",
				Password: "ValidPassword123456789",
			},
			setup: func(m *mocks.MockUserRepository) {
				hash, _ := security.HashPassword("ValidPassword123456789")
				m.EXPECT().LoginByUsername(gomock.Any(), &domain.LoginUserDb{Username: "ValidUsername"}).Return(&domain.LoginUserResponseDb{
					Id:           1,
					PasswordHash: hash,
				}, nil)
			},
			wantErr: nil,
		},
		{
			name: "IncorrectPasswordTest",
			input: domain.LoginUserInternal{
				Username: "ValidUsername",
				Password: "InvalidPassword123456789",
			},
			setup: func(m *mocks.MockUserRepository) {
				hash, _ := security.HashPassword("ValidPassword123456789")
				m.EXPECT().LoginByUsername(gomock.Any(), &domain.LoginUserDb{Username: "ValidUsername"}).Return(&domain.LoginUserResponseDb{
					Id:           1,
					PasswordHash: hash,
				}, nil)
			},
			wantErr: security.ErrIncorrectPassword,
		},
		{
			name: "dbErrLoginUser",
			input: domain.LoginUserInternal{
				Username: "ValidUsername",
				Password: "Password123456789",
			},
			setup: func(m *mocks.MockUserRepository) {
				m.EXPECT().LoginByUsername(gomock.Any(), &domain.LoginUserDb{Username: "ValidUsername"}).Return(nil, ErrInternalDb)
			},
			wantErr: ErrInternalDb,
		},
		{
			name: "UserNotFoundTest",
			input: domain.LoginUserInternal{
				Username: "NonExistentUser",
				Password: "ValidPassword123456789",
			},
			setup: func(m *mocks.MockUserRepository) {
				m.EXPECT().LoginByUsername(gomock.Any(), &domain.LoginUserDb{Username: "NonExistentUser"}).Return(nil, postgres.ErrUserNotFound)
			},
			wantErr: postgres.ErrUserNotFound,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mocks.NewMockUserRepository(ctrl)
			tc.setup(mockRepo)
			service := NewUserService(mockRepo)
			ctx := context.Background()
			res, err := service.UserLogin(ctx, &tc.input)

			if !errors.Is(err, tc.wantErr) {
				t.Errorf("UserLogin(%v) error = %v, wantErr %v", tc.input, err, tc.wantErr)
			}
			if tc.wantErr == nil {
				if res.JwtToken == "" {
					t.Errorf("UserLogin(%v) returned empty JWT token, expected non-empty", tc.input)
				}
			}
		})
	}
}
