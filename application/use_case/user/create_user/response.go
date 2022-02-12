package create_user

import (
	"app/models"
	"time"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	CreateUserResponse struct {
		base.BaseResponse
		Data CreateUserResponseData `json:"data"`
	}

	CreateUserResponseData struct {
		ID        int       `json:"id"`
		Username  string    `json:"username"`
		Role      int       `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func SetResponse(model CreateUserResponseData, message string, success bool) CreateUserResponse {
	return CreateUserResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: model,
	}
}

func ResponseMapper(model models.User) CreateUserResponseData {
	return CreateUserResponseData{
		ID:        model.ID,
		Username:  model.Username,
		Role:      model.Role,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
