package repository

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/domain"
	"context"
)

type AdsRepository interface {
	CreateAds(ctx context.Context, adsData *domain.CreateAdsDb) (*boundary.CreateAdsResponse, error)
	GetAdsList(ctx context.Context, adsListData *domain.AdsListDb) ([]*domain.AdsListResponseDb, int, error)
}
