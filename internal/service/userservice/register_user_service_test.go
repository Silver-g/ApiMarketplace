package userservice

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/domain"
	"ApiMarketplace/internal/mocks"
	"ApiMarketplace/internal/store/postgres"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

type registerTestCase struct {
	name    string
	input   domain.RegisterUserInternal
	setup   func(m *mocks.MockUserRepository)
	wantErr error
}

var ErrInternalDb = errors.New("db error")

func TestRegisterUser(t *testing.T) {
	cases := []registerTestCase{
		{
			name:  "TestUserAlreadyExists",
			input: domain.RegisterUserInternal{Username: "existing_user", Password: "somepass123456"},
			setup: func(m *mocks.MockUserRepository) {
				m.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil, postgres.ErrUserAlreadyExists)
			},
			wantErr: postgres.ErrUserAlreadyExists,
		},
		{
			name:  "dbErrRegisterUser",
			input: domain.RegisterUserInternal{Username: "new_user", Password: "somepass123456"},
			setup: func(m *mocks.MockUserRepository) {
				m.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil, ErrInternalDb)
			},
			wantErr: ErrInternalDb,
		},
		{
			name:  "TestValidRegister",
			input: domain.RegisterUserInternal{Username: "new_user", Password: "somepass123456"},
			setup: func(m *mocks.MockUserRepository) {
				m.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&boundary.RegisterUserResponse{Id: 1, Login: "new_user"}, nil)
			},
			wantErr: nil,
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
			_, err := service.Register(ctx, &tc.input)

			if !errors.Is(err, tc.wantErr) {
				t.Errorf("Register(%v) error = %v, wantErr %v", tc.input, err, tc.wantErr)
			}
		})
	}

}
