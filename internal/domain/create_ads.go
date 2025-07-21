package domain

import "github.com/shopspring/decimal"

type CreateAdsInternal struct {
	Id          int             `json:"id"`
	UserId      int             `json:"user_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ImageUrl    string          `json:"image_url"`
	Price       decimal.Decimal `json:"price"`
}
type CreateAdsDb struct {
	UserId      int             `json:"user_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ImageUrl    string          `json:"image_url"`
	Price       decimal.Decimal `json:"price"`
}
