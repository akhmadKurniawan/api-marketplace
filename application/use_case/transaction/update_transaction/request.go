package update_transaction

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type UpdateTransactionRequest struct {
	IdVa   string `json:"id_ve"`
	Status string `json:"status" db:"status"`
}

func ValidateRequest(req *UpdateTransactionRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RequestMapper(req UpdateTransactionRequest, status string) models.Transaction {
	return models.Transaction{
		IdVa:   req.IdVa,
		Status: status,
	}
}
