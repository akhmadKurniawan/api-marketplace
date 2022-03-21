package get_transaction

import "app/models"

type (
	ShowTransactionResponse struct {
		models.BaseResponse
		Data []ShowTransactionResponseData `json:"data"`
	}

	ShowTransactionResponseData struct {
		ID           int    `json:"id"`
		UserID       int    `json:"user_id"`
		Usser        User   `json:"user"`
		ProductID    int    `json:"product_id"`
		Type         string `json:"type"`
		Description  string `json:"description"`
		Amount       int    `json:"amount"`
		TotalProduct int    `json:"total_product"`
		Status       string `json:"status"`
	}

	User struct {
		UserID   int    `json:"user_id"`
		Username string `json:"user_name"`
	}

	Response struct {
		Transaction []models.Transaction
	}
)

func SetResponse(res *Response, msg string, scs bool) ShowTransactionResponse {
	return ShowTransactionResponse{
		BaseResponse: models.BaseResponse{
			Message: msg,
			Success: scs,
		},
		Data: ResponseMappers(res),
	}
}

func ResponseMappers(res *Response) []ShowTransactionResponseData {

	var list []ShowTransactionResponseData
	for _, val := range res.Transaction {
		response := ShowTransactionResponseData{
			ID:           val.ID,
			UserID:       val.UserID,
			ProductID:    val.ProductID,
			Type:         val.Type,
			Description:  val.Description,
			Amount:       val.Amount,
			TotalProduct: val.TotalProduct,
			Status:       val.Status,
		}
		list = append(list, response)
	}
	return list
}
