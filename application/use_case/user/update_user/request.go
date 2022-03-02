package update_user

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     int    `json:"role"`
}

func ValidateRequest(req *UpdateUserRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func RequestMapper(req UpdateUserRequest, username string, password string) models.User {
	return models.User{
		Username: username,
		Password: password,
		Role:     req.Role,
	}
}
