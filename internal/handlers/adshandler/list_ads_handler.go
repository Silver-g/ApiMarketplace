package adshandler

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/service/adsservice"
	"context"
	"net/http"
	"time"
)

type HandlerGetAdsList struct {
	Service adsservice.GetAdsListService
}

func NewHandlerGetAdsList(getAdsListServ adsservice.GetAdsListService) *HandlerGetAdsList {
	return &HandlerGetAdsList{Service: getAdsListServ}
}
func (h *HandlerGetAdsList) GetAdsListHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()
	if r.Method != http.MethodGet {
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only GET method is allowed",
		})
		return
	}
	userId, ok := r.Context().Value("user_id").(int)
	if !ok {
		boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
			ErrorCode: "InternalError",
			Message:   "Failed to get user Id from context",
		})
		return
	}

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sortBystr := r.URL.Query().Get("sort_by")
	priceFilterStr := r.URL.Query().Get("price")

	adsListQueryParamsMapping := boundary.AdsListQueryParamsMapping(pageStr, limitStr, sortBystr, priceFilterStr)
	adsListInternalMapping := boundary.ValidateAdsListQueryParams(adsListQueryParamsMapping, userId)
	responseData, err := h.Service.AdsList(ctx, &adsListInternalMapping)
	select {
	case <-ctx.Done():
		boundary.WriteResponseErr(w, 504, boundary.ErrorResponse{
			ErrorCode: "Timeout",
			Message:   "Request timed out",
		})
		return
	default:
		if err != nil {
			boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
				ErrorCode: "InternalError",
				Message:   err.Error(),
			})
			return
		}

		boundary.WriteResponseSuccess(w, 200, boundary.SuccessResponse{
			ResponseData: responseData,
			Message:      "Ads list sent successfully",
		})
	}
}
