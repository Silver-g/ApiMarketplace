package adshandler

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/service/adsservice"
	"context"
	"net/http"
	"time"
)

type HandlerCreateAds struct {
	Service adsservice.CreateAdsService
}

func NewHandlerCreateAds(createAdsServ adsservice.CreateAdsService) *HandlerCreateAds {
	return &HandlerCreateAds{Service: createAdsServ}
}

func (h *HandlerCreateAds) CreateAdsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()
	userId, ok := r.Context().Value("user_id").(int)
	if !ok {
		boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
			ErrorCode: "InternalError",
			Message:   "Failed to get user Id from context",
		})
		return
	}

	if r.Method != http.MethodPost {
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only POST method is allowed",
		})
		return
	}
	var adsCreateReq boundary.CreateAdsRequest
	err = boundary.DecodeJSONBody(r, &adsCreateReq)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   "Invalid request syntax",
		})
		return
	}
	err = boundary.ValidateCreateAdsRequest(adsCreateReq)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "ValidationError",
			Message:   err.Error(),
		})
		return
	}
	//
	createAdsMaping := boundary.CreateAdsMaping(adsCreateReq, userId)
	responseData, err := h.Service.CreateAds(ctx, &createAdsMaping)
	//
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

		boundary.WriteResponseSuccess(w, 201, boundary.SuccessResponse{
			ResponseData: responseData,
			Message:      "Ad successfully created",
		})
	}

}
