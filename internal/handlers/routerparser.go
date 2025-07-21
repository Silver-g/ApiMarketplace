package handlers

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/consts"
	"ApiMarketplace/internal/handlers/adshandler"
	"ApiMarketplace/internal/middleware"
	"net/http"
)

type RouteInfo struct {
	GetAdsListHandler *adshandler.HandlerGetAdsList
	CreateAdsHandler  *adshandler.HandlerCreateAds
}

func (ri *RouteInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	const prefix = "/ads"
	path := r.URL.Path
	//spliUrl := strings.TrimPrefix(r.URL.Path, prefix)
	// spliUrl = strings.Trim(spliUrl, "/")

	// var parts []string
	// if spliUrl != "" {
	// 	parts = strings.Split(spliUrl, "/")
	// }
	// раскоментить если нужно будет обработать сложные запросы /ads/1 ... и тд
	if path == prefix || path == prefix+"/" {
		if r.Method == http.MethodPost { // && len(parts) == 0 - если вдруг сложный url /ads/...
			handler := http.HandlerFunc(ri.CreateAdsHandler.CreateAdsHandler)
			middleware.AuthMiddleware(handler).ServeHTTP(w, r)
			return
		}
		if r.Method == http.MethodGet { // && len(parts) == 0
			handler := http.HandlerFunc(ri.GetAdsListHandler.GetAdsListHandler)
			middleware.OptionalAuthMiddleware(handler).ServeHTTP(w, r)
			return
		}
	}
	boundary.WriteResponseErr(w, 404, boundary.ErrorResponse{
		ErrorCode: "NotFound",
		Message:   consts.ErrInvalidRouteMsg,
	})
}
