package create_product_type

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type CreateProductTypeRequest struct {
	Name string `json:"name"`
}

func ValidateRequest(req *CreateProductTypeRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateProductTypeRequest) models.ProductType {
	return models.ProductType{
		Name: req.Name,
	}
}
