package create_product

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type CreateProductRequest struct {
	ProductType int    `bson:"product_type" json:"product_type" form:"product_type"`
	ShopId      int    `bson:"shop_id" json:"shop_id" form:"shop_id"`
	Name        string `bson:"name" json:"name" form:"name"`
	Price       int    `bson:"price" json:"price" form:"price"`
	Description string `bson:"description" json:"description" form:"description"`
	Qty         int    `bson:"qty" json:"qty" form:"qty"`
}

func ValidateRequest(req *CreateProductRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateProductRequest, img string) models.Product {
	return models.Product{
		ProductType: req.ProductType,
		ShopId:      req.ShopId,
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		Qty:         req.Qty,
		Image:       img,
	}
}
