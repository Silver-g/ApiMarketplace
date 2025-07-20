package adsservice

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/domain"
	"context"
	"fmt"
)

type GetAdsListService interface {
	AdsList(ctx context.Context, adsListData *domain.AdsListInternal) (*boundary.AdsListResponse, error)
}

func (s *AdsService) AdsList(ctx context.Context, adsListData *domain.AdsListInternal) (*boundary.AdsListResponse, error) {
	offset := (adsListData.Page - 1) * adsListData.Limit
	queryStr := BuildAdsListQuery(adsListData, offset)
	adsListDbMapData := boundary.AdsListDbMapping(*adsListData, queryStr)

	responseAdsListDbData, totalRecords, err := s.adsRepo.GetAdsList(ctx, &adsListDbMapData)
	if err != nil {
		return nil, err
	}
	totalPage := (totalRecords + adsListData.Limit - 1) / adsListData.Limit

	var items []boundary.AdsListItemResponse
	for _, ad := range responseAdsListDbData {
		isOwner := adsListData.UserId == ad.UserId
		item := boundary.AdsListItemResponseMapping(*ad, isOwner)
		items = append(items, item)
	}
	adsListResponse := boundary.AdsListResponseMapping(items, adsListData.Page, totalPage)
	return &adsListResponse, nil
}

func BuildAdsListQuery(params *domain.AdsListInternal, offset int) string {
	query := `SELECT ads.id, ads.user_id, users.login, ads.title, ads.description, ads.image_url, ads.price, COUNT(ads.id) OVER() AS total_records
	          FROM ads
	          JOIN users ON ads.user_id = users.id`
	if params.MinPrice != nil {
		query += fmt.Sprintf(" WHERE ads.price >= %d", *params.MinPrice)
	}
	if params.MaxPrice != nil {
		query += fmt.Sprintf(" AND ads.price <= %d", *params.MaxPrice)
	}
	query += fmt.Sprintf(" ORDER BY %s %s", params.SortType, params.SortOrder)
	if params.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", params.Limit)

	}
	if offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", offset)
	}
	return query
}
