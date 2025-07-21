package boundary

import (
	"strings"
	"testing"

	"github.com/shopspring/decimal"
)

type adsCase struct {
	name    string
	input   CreateAdsRequest
	wantErr error
}

func TestValidateCreateAdsRequest(t *testing.T) {
	cases := []adsCase{
		{
			name: "ValidInput",
			input: CreateAdsRequest{
				Title:       "Valid Ad",
				Description: "Valid description",
				Price:       decimal.NewFromFloat(99.99),
			},
			wantErr: nil,
		},
		{
			name: "EmptyTitle",
			input: CreateAdsRequest{
				Title:       "",
				Description: "Valid description",
				Price:       decimal.NewFromFloat(99.99),
			},
			wantErr: ErrEmptyAdsTitle,
		},
		{
			name: "OnlySpacesTitle",
			input: CreateAdsRequest{
				Title:       "   ",
				Description: "Valid description",
				Price:       decimal.NewFromFloat(99.99),
			},
			wantErr: ErrEmptyAdsTitle,
		},
		{
			name: "OnlyUnderscoresTitle",
			input: CreateAdsRequest{
				Title:       "___",
				Description: "Valid description",
				Price:       decimal.NewFromFloat(99.99),
			},
			wantErr: ErrEmptyAdsTitle,
		},
		{
			name: "EmptyDescription",
			input: CreateAdsRequest{
				Title:       "Valid Ad",
				Description: "",
				Price:       decimal.NewFromFloat(99.99),
			},
			wantErr: ErrEmptyAdsDescription,
		},
		{
			name: "OnlySpacesDescription",
			input: CreateAdsRequest{
				Title:       "Valid Ad",
				Description: "   ",
				Price:       decimal.NewFromFloat(99.99),
			},
			wantErr: ErrEmptyAdsDescription,
		},
		{
			name: "NegativePrice",
			input: CreateAdsRequest{
				Title:       "Valid Ad",
				Description: "Valid description",
				Price:       decimal.NewFromFloat(-10.00),
			},
			wantErr: ErrInvalidAdsPrice,
		},
		{
			name: "TooPrecisePrice",
			input: CreateAdsRequest{
				Title:       "Valid Ad",
				Description: "Valid description",
				Price:       decimal.NewFromFloat(99.999),
			},
			wantErr: ErrInvalidAdsPrice,
		},
		{
			name: "TooLargePrice",
			input: CreateAdsRequest{
				Title:       "Valid Ad",
				Description: "Valid description",
				Price:       decimal.NewFromInt(10000000000),
			},
			wantErr: ErrInvalidAdsPrice,
		},
		{
			name: "TooShortTitle",
			input: CreateAdsRequest{
				Title:       "A",
				Description: "Valid description",
				Price:       decimal.NewFromFloat(99.99),
			},
			wantErr: ErrInvalidAdsTitle,
		},
		{
			name: "TooLongTitle",
			input: CreateAdsRequest{
				Title:       longString(),
				Description: "Valid description",
				Price:       decimal.NewFromFloat(99.99),
			},
			wantErr: ErrInvalidAdsTitle,
		},
		{
			name: "TooLongDescription",
			input: CreateAdsRequest{
				Title:       "Valid Ad",
				Description: strings.Repeat("a", 8001),
				Price:       decimal.NewFromFloat(99.99),
			},
			wantErr: ErrInvalidAdsDescription,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateCreateAdsRequest(tc.input)
			if err != tc.wantErr {
				t.Errorf("ValidateCreateAdsRequest error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
func longString() string {
	str := ""
	for i := 0; i < 151; i++ {
		str += "a"
	}
	return str
}
