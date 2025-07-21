package domain

import "github.com/shopspring/decimal"

type AdsListInternal struct {
	Page        int
	Limit       int
	SortType    string // price, created_at
	SortOrder   string // ASC или DESC
	MinPrice    *int
	MaxPrice    *int
	Id          int
	UserId      int
	Title       string
	Description string
	ImageUrl    string
	Price       decimal.Decimal
}

type AdsListDb struct {
	QueryString string
	Id          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ImageUrl    string          `json:"image_url"`
	Price       decimal.Decimal `json:"price"`
	AuthorLogin string          `json:"author_login"`
	UserId      int             `json:"user_id"`
}

type AdsListResponseDb struct {
	Id          int             `json:"id"`
	UserId      int             `json:"user_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ImageUrl    string          `json:"image_url"`
	Price       decimal.Decimal `json:"price"`
	AuthorLogin string          `json:"author_login"`
}
