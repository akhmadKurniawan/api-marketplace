package create_product

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type CreateProductRequest struct {
	ProductType int    `json:"product_type" form:"product_type"`
	ShopId      int    `json:"shop_id" form:"shop_id"`
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Qty         string `json:"qty" form:"qty"`
	UserID      int
}

func ValidateRequest(req *CreateProductRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateProductRequest) models.Product {
	return models.Product{
		ProductType: req.ProductType,
		ShopId:      req.ShopId,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Qty:         req.Qty,
	}
}
