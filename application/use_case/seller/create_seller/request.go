package create_seller

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type CreateSellerRequest struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name" validate:"required"`
	Alamat string `json:"alamat" validate:"required"`
	NoHp   int    `json:"no_hp" validate:"required"`
}

func ValidateRequest(req *CreateSellerRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateSellerRequest) models.Seller {
	return models.Seller{
		UserID: req.UserID,
		Name:   req.Name,
		Alamat: req.Alamat,
		NoHp:   req.NoHp,
	}
}
