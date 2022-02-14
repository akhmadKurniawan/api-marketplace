package delete_seller

import (
	base "github.com/refactory-id/go-core-package/response"
)

type DeleteSellerResponse struct {
	base.BaseResponse
}

func SetResponse(message string, success bool) DeleteSellerResponse {
	return DeleteSellerResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
	}
}
