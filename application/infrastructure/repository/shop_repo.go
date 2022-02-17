package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"

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

func (repo *ShopRepository) CreateShop(ctx context.Context, shop models.Shop, ImgFile string) error {
	db := repo.DB

	shop.Logo = ImgFile

	errCreate := db.Create(&shop).Error
	if errCreate != nil {
		return errCreate
	}

	return nil
}
