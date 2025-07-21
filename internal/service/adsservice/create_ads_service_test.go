package adsservice

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/domain"
	"ApiMarketplace/internal/mocks"
	"ApiMarketplace/internal/store/postgres"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
)

type createAdsTestCase struct {
	name    string
	input   domain.CreateAdsInternal
	setup   func(m *mocks.MockAdsRepository)
	wantErr error
}

func TestCreateAds(t *testing.T) {
	cases := []createAdsTestCase{
		{
			name: "TestUserNoFound",
			input: domain.CreateAdsInternal{
				UserId:      123,
				Id:          1,
				Title:       "TestTitle",
				Description: "Test description",
				ImageUrl:    "https/URL",
				Price:       decimal.NewFromFloat(99.99),
			},
			setup: func(m *mocks.MockAdsRepository) {
				m.EXPECT().CreateAds(gomock.Any(), gomock.Any()).Return(nil, postgres.ErrUserNotFound)
			},
			wantErr: postgres.ErrUserNotFound,
		},
		{
			name: "dbErrCreateAds",
			input: domain.CreateAdsInternal{
				UserId:      123,
				Id:          1,
				Title:       "TestTitle",
				Description: "Test description",
				ImageUrl:    "https/URL",
				Price:       decimal.NewFromFloat(99.99),
			},
			setup: func(m *mocks.MockAdsRepository) {
				m.EXPECT().CreateAds(gomock.Any(), gomock.Any()).Return(nil, ErrInternalDb)
			},
			wantErr: ErrInternalDb,
		},
		{
			name: "TestValidCreateAds",
			input: domain.CreateAdsInternal{
				UserId:      1,
				Id:          1,
				Title:       "TestTitle",
				Description: "Test description",
				ImageUrl:    "https/URL",
				Price:       decimal.NewFromFloat(99.99),
			},
			setup: func(m *mocks.MockAdsRepository) {
				m.EXPECT().CreateAds(gomock.Any(), gomock.Any()).Return(&boundary.CreateAdsResponse{
					Id:          1,
					Title:       "TestTitle",
					Description: "Test description",
					ImageUrl:    "https/URL",
					Price:       decimal.NewFromFloat(99.99),
				}, nil)
			},
			wantErr: nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := mocks.NewMockAdsRepository(ctrl)
			tc.setup(mockRepo)
			service := NewAdsService(mockRepo)
			ctx := context.Background()
			_, err := service.CreateAds(ctx, &tc.input)

			if !errors.Is(err, tc.wantErr) {
				t.Errorf("Register(%v) error = %v, wantErr %v", tc.input, err, tc.wantErr)
			}
		})
	}
}
