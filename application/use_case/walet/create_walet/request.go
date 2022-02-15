package create_walet

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type CreateWaletRequest struct {
	UserID int `json:"user_id"`
	Saldo  int `json:"saldo"`
}

func ValidateRequest(req *CreateWaletRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateWaletRequest, userID int) models.Walet {
	return models.Walet{
		UserID: userID,
		Saldo:  req.Saldo,
	}
}
