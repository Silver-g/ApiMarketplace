package adsservice

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/domain"
	"ApiMarketplace/internal/store/repository"
	"context"
)

type CreateAdsService interface {
	CreateAds(ctx context.Context, adsData *domain.CreateAdsInternal) (*boundary.CreateAdsResponse, error)
}
type AdsService struct {
	adsRepo repository.AdsRepository
}

func NewAdsService(repo repository.AdsRepository) *AdsService {
	return &AdsService{adsRepo: repo}
}
func (s *AdsService) CreateAds(ctx context.Context, adsData *domain.CreateAdsInternal) (*boundary.CreateAdsResponse, error) {
	adsDbData := boundary.CreateAdsDbMaping(*adsData)
	responseAdsData, err := s.adsRepo.CreateAds(ctx, &adsDbData)
	if err != nil {
		return nil, err
	}
	return responseAdsData, nil
}
