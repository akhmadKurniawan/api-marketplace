package create_product

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type CreateProductRequest struct {
	ProductType int    `json:"product_type"`
	ShopId      int    `json:"shop_id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Qty         string `json:"qty"`
	Image       string `json:"image"`
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

func RequestMapper(req CreateProductRequest, ShopId int) models.Product {
	return models.Product{
		ProductType: req.ProductType,
		ShopId:      ShopId,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Qty:         req.Qty,
		Image:       req.Image,
	}
}
