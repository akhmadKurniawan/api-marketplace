package create_transaction

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type CreateTransactionRequest struct {
	UserID       int    `json:"user_id"`
	ProductID    int    `json:"product_id"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	Amount       int    `json:"amount"`
	TotalProduct int    `json:"total_product"`
}

func ValidateRequest(req *CreateTransactionRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateTransactionRequest, amount int, typeE string, status string) models.Transaction {
	return models.Transaction{
		UserID:       req.UserID,
		ProductID:    req.ProductID,
		Type:         typeE,
		Description:  req.Description,
		Status:       status,
		Amount:       amount,
		TotalProduct: req.TotalProduct,
	}
}
