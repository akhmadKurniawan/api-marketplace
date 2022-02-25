package create_transaction

import base "github.com/refactory-id/go-core-package/response"

type (
	CreateTransactionResponse struct {
		base.BaseResponse
		Data CreateTransactionResponseData
	}

	CreateTransactionResponseData struct {
		Payment string `json:"payment"`
	}
)

func SetResponse(data CreateTransactionResponseData, message string, success bool) CreateTransactionResponse {
	return CreateTransactionResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: data,
	}
}

func ResponseMapper(payment string) CreateTransactionResponseData {
	return CreateTransactionResponseData{
		Payment: payment,
	}
}
