package update_user

import (
	"app/models"
	"time"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateUserResponse struct {
		base.BaseResponse
		Data UpdateUserResponseData `json:"data"`
	}

	UpdateUserResponseData struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Username  string    `json:"username"`
		Role      int       `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Response struct {
		User models.User
	}
)

func SetResponse(domain *Response, message string, success bool) UpdateUserResponse {
	return UpdateUserResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: ResponseMapper(domain),
	}
}

func ResponseMapper(domain *Response) UpdateUserResponseData {
	return UpdateUserResponseData{
		ID:        domain.User.Model.ID,
		Username:  domain.User.Username,
		Role:      domain.User.Role,
		CreatedAt: domain.User.Model.CreatedAt,
		UpdatedAt: domain.User.Model.UpdatedAt,
	}
}
