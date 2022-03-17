package login

import (
	"app/models"
	"time"
)

type (
	LoginResponse struct {
		models.BaseResponse
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

func SetResponse(res *Response, message string, success bool) LoginResponse {
	return LoginResponse{
		BaseResponse: models.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: ResponseMapper(res),
	}
}

func ResponseMapper(res *Response) LoginResponseData {
	return LoginResponseData{
		UserID:    res.User.ID,
		Token:     res.User.UserToken.Token,
		Username:  res.User.Username,
		Role:      res.User.Role,
		CreatedAt: res.User.CreatedAt,
		UpdatedAt: res.User.UpdatedAt,
	}
}
