package postgres

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/domain"
	"context"
	"database/sql"
)

type AdsPostgres struct {
	db *sql.DB
}

func NewAdsPostgres(dataBase *sql.DB) *AdsPostgres {
	return &AdsPostgres{db: dataBase}
}

func (r *AdsPostgres) GetAdsList(ctx context.Context, adsListData *domain.AdsListDb) ([]*domain.AdsListResponseDb, int, error) {
	var adsListItems []*domain.AdsListResponseDb
	var totalRecords int
	query := adsListData.QueryString
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	for rows.Next() {
		var ad domain.AdsListResponseDb
		err := rows.Scan(
			&ad.Id, &ad.UserId, &ad.AuthorLogin,
			&ad.Title, &ad.Description, &ad.ImageUrl, &ad.Price, &totalRecords,
		)
		if err != nil {
			return nil, 0, err
		}
		adsListItems = append(adsListItems, &ad)
	}
	return adsListItems, totalRecords, nil
}

func (r *AdsPostgres) CreateAds(ctx context.Context, adsData *domain.CreateAdsDb) (*boundary.CreateAdsResponse, error) {
	var adsId int
	query := "INSERT INTO ads (user_id, title, description, image_url, price) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, adsData.UserId, adsData.Title, adsData.Description, adsData.ImageUrl, adsData.Price).Scan(&adsId)
	if err != nil {
		return nil, err
	}
	respData := boundary.CreateAdsResponseMaping(*adsData, adsId)
	return &respData, nil
}
