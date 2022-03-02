package update_transaction

import (
	"app/models"
)

type (
	UpdateTransactionResponseData struct {
		IdVa    string `json:"external_id"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	UpdateTransactionResponse struct {
		Data UpdateTransactionResponseData
	}

	Response struct {
		Trans     models.Transaction
		Status    string `json:"status"`
		Message   string `json:"message"`
		ErrorCode string `json:"error_code"`
	}
)

func SetResponse(model *Response) UpdateTransactionResponse {
	return UpdateTransactionResponse{
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
