package create_costumer

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type CreateCostumerRequest struct {
	UserID string `json:"user_id"`
	Name   string `json:"name" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	NoHp   string `json:"no_hp" validate:"required"`
}

func ValidateRequest(req *CreateCostumerRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateCostumerRequest, userId int) models.Costumer {
	return models.Costumer{
		UserID: userId,
		Name:   req.Name,
		Alamat: req.Alamat,
		NoHp:   req.NoHp,
	}
}
