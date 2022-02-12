package create_user

import (
	"app/models"

	validator "github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     int    `json:"role"`
}

func ValidateRequest(req *CreateUserRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

var _default int = 1

func RequestMapper(req CreateUserRequest, password string) models.User {
	return models.User{
		Username: req.Username,
		Password: password,
		Role:     _default,
	}
}
