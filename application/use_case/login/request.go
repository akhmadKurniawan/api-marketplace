package login

import (
	"app/models"

	validator "github.com/go-playground/validator/v10"
)

type (
	LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		UserID   int    `json:"user_id"`
		Token    string `json:"token"`
	}
)

func ValidateRequest(req *LoginRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(userId int, signed string) models.UserToken {
	return models.UserToken{
		UserID: userId,
		Token:  signed,
	}
}
