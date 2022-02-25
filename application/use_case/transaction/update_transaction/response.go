package update_transaction

import (
	"app/models"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateTransactionResponseData struct {
		IdVa    string `json:"external_id"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	UpdateTransactionResponse struct {
		base.BaseResponse
		Data UpdateTransactionResponseData
	}

	Response struct {
		Trans     models.Transaction
		Status    string `json:"status"`
		Message   string `json:"message"`
		ErrorCode string `json:"error_code"`
	}
)

func SetResponse(model *Response, message string, success bool) UpdateTransactionResponse {
	return UpdateTransactionResponse{
		BaseResponse: base.BaseResponse{
			Message: message,
			Success: success,
		},
		Data: ResponseMapper(model),
	}
}

func ResponseMapper(model *Response) UpdateTransactionResponseData {
	return UpdateTransactionResponseData{
		IdVa:    model.Trans.IdVa,
		Status:  model.Trans.Status,
		Message: model.Trans.Message,
	}
}
