package create_shop

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type CreateShopRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Alamat      string `json:"alamat" form:"alamat"`
	UserID      int
}

func ValidateRequest(req *CreateShopRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMapper(req CreateShopRequest, sellerID int, logo string) models.Shop {
	return models.Shop{
		Name:        req.Name,
		Description: req.Description,
		Alamat:      req.Alamat,
		Logo:        logo,
		SellerID:    sellerID,
	}
}
