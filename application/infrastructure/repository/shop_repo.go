package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type ShopRepository struct {
	DB *gorm.DB
}

func NewShopRepository(db *gorm.DB) infrastructure.ShopRepository {
	return &ShopRepository{
		DB: db,
	}
}

func (repo *ShopRepository) CreateShop(ctx context.Context, shop models.Shop) error {
	db := repo.DB

	errCreate := db.Create(&shop).Error
	if errCreate != nil {
		return errCreate
	}

	return nil
}

func (repo *ShopRepository) GetShopBySellerID(ctx context.Context, id int) (models.Shop, error) {
	shop := models.Shop{}

	if err := repo.DB.Where("seller_id = ?", id).First(&shop).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Repository - GetShopBySellerID Error : ", err)
		}
		return shop, err
	}

	return shop, nil
}
