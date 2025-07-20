package userhandler

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/service/userservice"
	"ApiMarketplace/internal/store/postgres"
	"context"
	"net/http"
	"time"
)

type HandlerRegisterUser struct {
	Service userservice.UserRegister
}

func NewHandlerRegister(regServ userservice.UserRegister) *HandlerRegisterUser {
	return &HandlerRegisterUser{Service: regServ}
}

func (h *HandlerRegisterUser) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	//  контекст с таймаутом 20 секунд
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()

	if r.Method != http.MethodPost {
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only POST method is allowed.",
		})
		return
	}

	var userReq boundary.UserRequest
	err := boundary.DecodeJSONBody(r, &userReq)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "StatusBadRequest",
			Message:   "Invalid syntax",
		})
		return
	}

	err = boundary.UserValidate(userReq)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "ValidationError",
			Message:   err.Error(),
		})
		return
	}

	regUserMaping := boundary.RegisterUserMaping(userReq)
	responseData, err := h.Service.Register(ctx, &regUserMaping)

	select {
	case <-ctx.Done():
		boundary.WriteResponseErr(w, 504, boundary.ErrorResponse{
			ErrorCode: "Timeout",
			Message:   "Request timed out",
		})
		return
	default:
		if err != nil {
			if err == postgres.ErrUserAlreadyExists {
				boundary.WriteResponseErr(w, 409, boundary.ErrorResponse{
					ErrorCode: "UserAlreadyExists",
					Message:   "Conflict: User with this username already exists.",
				})
				return
			}
			boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
				ErrorCode: "InternalError",
				Message:   err.Error(),
			})
			return
		}

		boundary.WriteResponseSuccess(w, 201, boundary.SuccessResponse{
			ResponseData: responseData,
			Message:      "User successfully created",
		})
	}
}
