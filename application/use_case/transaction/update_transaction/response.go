package update_transaction

import (
	"app/models"

	base "github.com/refactory-id/go-core-package/response"
)

type (
	UpdateTransactionResponseData struct {
		IdVa   string `json:"id_va"`
		Status string `json:"status"`
		// ID           int       `json:"id"`
		// UserId       int       `json:"user_id"`
		// ProductId    int       `json:"product_id"`
		// Description  string    `json:"description"`
		// Amount       int       `json:"amount"`
		// TotalProduct int       `json:"total_product"`
		// CreatedAt    time.Time `json:"created_at"`
		// UpdatedAt    time.Time `json:"updated_at"`
	}
	UpdateTransactionResponse struct {
		base.BaseResponse
		Data UpdateTransactionResponseData
	}

	Response struct {
		Trans models.Transaction
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
		IdVa:   model.Trans.IdVa,
		Status: model.Trans.Status,
		// ID:           model.Trans.Model.ID,
		// UserId:       model.Trans.UserID,
		// ProductId:    model.Trans.ProductID,
		// Description:  model.Trans.Description,
		// Amount:       model.Trans.Amount,
		// TotalProduct: model.Trans.TotalProduct,
		// CreatedAt:    model.Trans.CreatedAt,
		// UpdatedAt:    model.Trans.UpdatedAt,
	}
}
