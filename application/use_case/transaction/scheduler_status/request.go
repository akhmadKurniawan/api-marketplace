package scheduler_status

import (
	"app/models"
	"time"

	"github.com/go-playground/validator/v10"
)

type UpdateSchedulerRequest struct {
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ValidateRequest(req *UpdateSchedulerRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RequestMapper(req UpdateSchedulerRequest) models.Transaction {
	return models.Transaction{
		Status:    req.Status,
		UpdatedAt: req.UpdatedAt,
	}
}
