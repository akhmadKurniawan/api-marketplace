package verify_email_user

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type VerifyEmailUserRequest struct {
	Status string `json:"status"`
}

func ValidateRequest(req *VerifyEmailUserRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RequestMapper(req VerifyEmailUserRequest) models.User {
	return models.User{
		Status: req.Status,
	}
}
