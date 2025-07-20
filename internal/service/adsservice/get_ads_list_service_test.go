package adsservice

import (
	"ApiMarketplace/internal/domain"
	"ApiMarketplace/internal/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
)

var ErrInternalDb = errors.New("db error")

type getAdsListTestCase struct {
	name    string
	input   domain.AdsListInternal
	setup   func(m *mocks.MockAdsRepository)
	wantErr error
}

func ptr[T any](v T) *T {
	return &v
}

func TestGetAdsList(t *testing.T) {
	cases := []getAdsListTestCase{
		{
			name: "SuccessGetListAds",
			input: domain.AdsListInternal{
				Page:        1,
				Limit:       10,
				SortType:    "ads.price",
				SortOrder:   "ASC",
				MinPrice:    ptr(100),
				MaxPrice:    ptr(2000),
				Id:          3,
				UserId:      5,
				Title:       "TestOk",
				Description: "Test decription",
				ImageUrl:    "https/URL",
				Price:       decimal.NewFromFloat(99.99),
			},
			setup: func(m *mocks.MockAdsRepository) {
				m.EXPECT().GetAdsList(gomock.Any(), gomock.Any()).Return([]*domain.AdsListResponseDb{
					{
						Id:          3,
						UserId:      5,
						Title:       "TestOk",
						Description: "description OK",
						ImageUrl:    "https/URL",
						Price:       decimal.NewFromFloat(99.99),
						AuthorLogin: "LoginUser",
					},
				}, 1, nil)
			},
			wantErr: nil,
		},
		{
			name: "dbErrGetListAds",
			input: domain.AdsListInternal{
				Page:        2,
				Limit:       10,
				SortType:    "ads.price",
				SortOrder:   "ASC",
				MinPrice:    ptr(100),
				MaxPrice:    ptr(2000),
				Id:          4,
				UserId:      6,
				Title:       "TestOk",
				Description: "Test decription",
				ImageUrl:    "https/URL",
				Price:       decimal.NewFromFloat(99.99),
			},
			setup: func(m *mocks.MockAdsRepository) {
				m.EXPECT().GetAdsList(gomock.Any(), gomock.Any()).Return(nil, 0, ErrInternalDb)
			},
			wantErr: ErrInternalDb,
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
			_, err := service.AdsList(ctx, &tc.input)

			if !errors.Is(err, tc.wantErr) {
				t.Errorf("Register(%v) error = %v, wantErr %v", tc.input, err, tc.wantErr)
			}
		})
	}
}
