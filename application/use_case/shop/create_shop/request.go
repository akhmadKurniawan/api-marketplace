package create_shop

import (
	"app/models"

	"github.com/go-playground/validator/v10"
)

type CreateShopRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Alamat      string `json:"alamat"`
	Logo        string `json:"logo"`
	SellerID    int    `json:"seller_id"`
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

func RequestMapper(req CreateShopRequest, sellerID int) models.Shop {
	return models.Shop{
		Name:        req.Name,
		Description: req.Description,
		Alamat:      req.Alamat,
		Logo:        req.Logo,
		SellerID:    sellerID,
	}
}
