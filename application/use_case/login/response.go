package login

import (
	"app/models"
	"time"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	LoginResponse struct {
		base.BaseResponse
		Data LoginResponseData `json:"data"`
	}

	LoginResponseData struct {
		UserID    int       `json:"user_id"`
		Token     string    `json:"token"`
		Username  string    `json:"username"`
		Role      int       `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Response struct {
		User models.User
	}
)

func SetResponse(models *Response, message string, success bool) LoginResponse {
	return LoginResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: ResponseMapper(models),
	}
}

func ResponseMapper(models *Response) LoginResponseData {
	return LoginResponseData{
		UserID:    models.User.ID,
		Token:     models.User.UserToken.Token,
		Username:  models.User.Username,
		Role:      models.User.Role,
		CreatedAt: models.User.CreatedAt,
		UpdatedAt: models.User.UpdatedAt,
	}
}
