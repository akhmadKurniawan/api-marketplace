package create_user

import (
	"app/models"

	validator "github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     int    `json:"role" gorm:"default:1"`
	Name     string `json:"name"`
	Alamat   string `json:"alamat"`
	NoHp     string `json:"no_hp"`
}

func ValidateRequest(req *CreateUserRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func RequestMappers(req CreateUserRequest, id int) (models.Seller, models.Costumer) {
	reqSeller := models.Seller{
		UserID: id,
		Name:   req.Name,
		Alamat: req.Alamat,
		NoHp:   req.NoHp,
	}

	reqCostumer := models.Costumer{
		UserID: id,
		Name:   req.Name,
		Alamat: req.Alamat,
		NoHp:   req.NoHp,
	}
	return reqSeller, reqCostumer
}

func RequestMapper(req CreateUserRequest, password string) models.User {
	reqUser := models.User{
		Username: req.Username,
		Password: password,
		Role:     req.Role,
	}

	// reqSeller := models.Seller{
	// 	UserID: id,
	// 	Name:   req.Name,
	// 	Alamat: req.Alamat,
	// 	NoHp:   req.NoHp,
	// }

	// reqCostumer := models.Costumer{
	// 	UserID: id,
	// 	Name:   req.Name,
	// 	Alamat: req.Alamat,
	// 	NoHp:   req.NoHp,
	// }

	return reqUser
}
