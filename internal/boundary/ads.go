package boundary

import (
	"ApiMarketplace/internal/consts"
	"ApiMarketplace/internal/domain"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

type AdsListQueryParamsRequest struct {
	PageStr  string `json:"page"`
	LimitStr string `json:"limit"`
	SortBy   string `json:"sort_by"`
	PriceStr string `json:"sort_order"`
}

type AdsListItemResponse struct {
	Id          int             `json:"id"`
	UserName    string          `json:"login"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ImageUrl    string          `json:"image_url"`
	Price       decimal.Decimal `json:"price"`
	IsOwner     bool            `json:"is_owner"`
}

type AdsListResponse struct {
	Ads        []*AdsListItemResponse `json:"ads"`
	Page       int                    `json:"page"`
	TotalPages int                    `json:"total_pages"`
}

type CreateAdsRequest struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ImageUrl    string          `json:"image_url"`
	Price       decimal.Decimal `json:"price"`
}

type CreateAdsResponse struct {
	Id          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ImageUrl    string          `json:"image_url"`
	Price       decimal.Decimal `json:"price"`
}

var (
	ErrEmptyAdsTitle         = errors.New(consts.ErrEmptyAdsTitleMsg)
	ErrEmptyAdsDescription   = errors.New(consts.ErrEmptyAdsDescriptionMsg)
	ErrInvalidAdsPrice       = errors.New(consts.ErrInvalidAdsPriceMsg)
	ErrInvalidAdsTitle       = errors.New(consts.ErrInvalidAdsTitleMsg)
	ErrInvalidAdsDescription = errors.New(consts.ErrInvalidAdsDescriptionMsg)
)

func CreateAdsMaping(createAdsReq CreateAdsRequest, userId int) domain.CreateAdsInternal {
	return domain.CreateAdsInternal{
		UserId:      userId,
		Title:       createAdsReq.Title,
		Description: createAdsReq.Description,
		ImageUrl:    createAdsReq.ImageUrl,
		Price:       createAdsReq.Price,
	}
}

func CreateAdsDbMaping(createAdsReq domain.CreateAdsInternal) domain.CreateAdsDb {
	return domain.CreateAdsDb{
		UserId:      createAdsReq.UserId,
		Title:       createAdsReq.Title,
		Description: createAdsReq.Description,
		ImageUrl:    createAdsReq.ImageUrl,
		Price:       createAdsReq.Price,
	}
}

func CreateAdsResponseMaping(createAdsReq domain.CreateAdsDb, adsId int) CreateAdsResponse {
	return CreateAdsResponse{
		Id:          adsId,
		Title:       createAdsReq.Title,
		Description: createAdsReq.Description,
		ImageUrl:    createAdsReq.ImageUrl,
		Price:       createAdsReq.Price,
	}
}

func AdsListItemResponseMapping(adsListDbData domain.AdsListResponseDb, ownerFlag bool) *AdsListItemResponse {
	return &AdsListItemResponse{
		Id:          adsListDbData.Id,
		UserName:    adsListDbData.AuthorLogin,
		Title:       adsListDbData.Title,
		Description: adsListDbData.Description,
		ImageUrl:    adsListDbData.ImageUrl,
		Price:       adsListDbData.Price,
		IsOwner:     ownerFlag,
	}
}

func AdsListResponseMapping(items []*AdsListItemResponse, page int, totalPage int) *AdsListResponse {
	return &AdsListResponse{
		Ads:        items,
		Page:       page,
		TotalPages: totalPage,
	}
}

func AdsListQueryParamsMapping(pageStr string, limitStr string, sortBy string, priceStr string) AdsListQueryParamsRequest {
	return AdsListQueryParamsRequest{
		PageStr:  pageStr,
		LimitStr: limitStr,
		SortBy:   sortBy,
		PriceStr: priceStr,
	}
}

func AdsListInternalMapping(userId int, page int, limit int, sortType string, sortOrder string, min *int, max *int) domain.AdsListInternal {
	return domain.AdsListInternal{
		UserId:    userId,
		Page:      page,
		Limit:     limit,
		SortType:  sortType,
		SortOrder: sortOrder,
		MinPrice:  min,
		MaxPrice:  max,
	}
}

func AdsListDbMapping(adsListData domain.AdsListInternal, queryStr string) domain.AdsListDb {
	return domain.AdsListDb{
		QueryString: queryStr,
		Id:          adsListData.Id,
		Title:       adsListData.Title,
		Description: adsListData.Description,
		ImageUrl:    adsListData.ImageUrl,
		Price:       adsListData.Price,
		UserId:      adsListData.UserId,
	}
}

const (
	defaultPage     = 1
	defaultLimit    = 10
	MaxAllowedPrice = 10_000_000_000
)

// допустимые значения для sort_by и их маппинг на SortField и SortOrder
var validSortBy = map[string]struct {
	SortType  string
	SortOrder string
}{
	"price_low":      {SortType: "ads.price", SortOrder: "ASC"},
	"price_high":     {SortType: "ads.price", SortOrder: "DESC"},
	"created_at_new": {SortType: "ads.created_at", SortOrder: "DESC"},
	"created_at_old": {SortType: "ads.created_at", SortOrder: "ASC"},
}

func ValidateAdsListQueryParams(req AdsListQueryParamsRequest, user_id int) domain.AdsListInternal {

	page, err := strconv.Atoi(req.PageStr)
	if err != nil || page < 1 {
		page = defaultPage
	}

	limit, err := strconv.Atoi(req.LimitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = defaultLimit
	}

	//  sort_by
	sortbyPars, ok := validSortBy[req.SortBy]
	if !ok {
		sortbyPars = validSortBy["created_at_new"] // дефолт
	}

	// price filter
	var minPrice *int
	var maxPrice *int
	if req.PriceStr != "" {
		parts := strings.Split(req.PriceStr, "-")
		if len(parts) == 2 {
			min, err1 := strconv.Atoi(parts[0])
			max, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil && min >= 0 && max <= MaxAllowedPrice && min <= max {
				minPrice = &min
				maxPrice = &max
			} else {
				fmt.Println("Некорректные значения фильтра цены:", req.PriceStr) // логи с проблемой
			}
		} else {
			fmt.Println("Неверный формат фильтра цены:", req.PriceStr) // логи с проблемой
		}
	}
	result := AdsListInternalMapping(user_id, page, limit, sortbyPars.SortType, sortbyPars.SortOrder, minPrice, maxPrice)
	return result
}

func ValidateCreateAdsRequest(adsReq CreateAdsRequest) error {
	if strings.TrimSpace(adsReq.Title) == "" || isOnlySpacesOrUnderscores(adsReq.Title) { //намеренно не задавал регулярное выражение
		return ErrEmptyAdsTitle
	}
	if strings.TrimSpace(adsReq.Description) == "" || isOnlySpacesOrUnderscores(adsReq.Description) {
		return ErrEmptyAdsDescription
	}
	if adsReq.Price.IsNegative() {
		return ErrInvalidAdsPrice
	}
	if adsReq.Price.Exponent() < -2 {
		return ErrInvalidAdsPrice
	}
	if adsReq.Price.GreaterThan(decimal.NewFromInt(9999999999)) {
		return ErrInvalidAdsPrice
	}
	if len(adsReq.Title) < 2 || len(adsReq.Title) > 150 {
		return ErrInvalidAdsTitle
	}
	if len(adsReq.Description) > 8000 {
		return ErrInvalidAdsDescription
	}

	return nil
}
