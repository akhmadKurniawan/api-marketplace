package update_transaction

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type UpdateTransactionRequest struct {
	IdVa   string `json:"external_id"`
	Status string `json:"status" db:"status"`
	Amount int    `json:"amount"`
}

func ValidateRequest(req *UpdateTransactionRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RequestMapper(req UpdateTransactionRequest, id, message, status string) models.Transaction {
	return models.Transaction{
		IdVa:    id,
		Status:  status,
		Message: message,
		Amount:  req.Amount,
	}
}
