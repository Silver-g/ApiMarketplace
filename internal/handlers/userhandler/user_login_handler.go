package userhandler

import (
	"ApiMarketplace/internal/boundary"
	"ApiMarketplace/internal/security"
	"ApiMarketplace/internal/service/userservice"
	"ApiMarketplace/internal/store/postgres"
	"context"
	"errors"
	"net/http"
	"time"
)

type HandlerLoginUser struct {
	Service userservice.UserLogin
}

func NewLoginHandler(loginServ userservice.UserLogin) *HandlerLoginUser {
	return &HandlerLoginUser{Service: loginServ}
}

func (h *HandlerLoginUser) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
	defer cancel()
	if r.Method != http.MethodPost {
		boundary.WriteResponseErr(w, 405, boundary.ErrorResponse{
			ErrorCode: "MethodNotAllowed",
			Message:   "Only POST method is allowed",
		})
		return
	}
	var userReq boundary.UserRequest
	err = boundary.DecodeJSONBody(r, &userReq)
	if err != nil {
		boundary.WriteResponseErr(w, 400, boundary.ErrorResponse{
			ErrorCode: "BadRequest",
			Message:   "Invalid request syntax",
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
	loginUserMaping := boundary.LoginUserMaping(userReq)
	responseData, err := h.Service.UserLogin(ctx, &loginUserMaping)

	select {
	case <-ctx.Done():
		boundary.WriteResponseErr(w, 504, boundary.ErrorResponse{
			ErrorCode: "Timeout",
			Message:   "Request timed out",
		})
		return
	default:
		if err != nil {
			if errors.Is(err, postgres.ErrUserNotFound) {
				boundary.WriteResponseErr(w, 404, boundary.ErrorResponse{
					ErrorCode: "UserNotFound",
					Message:   "User with this username not found",
				})
				return
			}
			if err == security.ErrIncorrectPassword {
				boundary.WriteResponseErr(w, 401, boundary.ErrorResponse{
					ErrorCode: "IncorrectPassword",
					Message:   err.Error(),
				})
				return
			}
			boundary.WriteResponseErr(w, 500, boundary.ErrorResponse{
				ErrorCode: "InternalError",
				Message:   err.Error(),
			})
			return
		}

		boundary.WriteResponseSuccess(w, 200, boundary.SuccessResponse{
			ResponseData: responseData,
			Message:      "User successfully logged in",
		})
	}
}
