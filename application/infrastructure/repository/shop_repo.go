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

func (repo *ShopRepository) CreateShop(ctx context.Context, shop models.Shop) error {
	db := repo.DB

	errCreate := db.Create(&shop).Error
	if errCreate != nil {
		return errCreate
	}

	return nil
}

func (repo *ShopRepository) GetShopById(ctx context.Context, id int) (models.Shop, error) {
	db := repo.DB
	shop := models.Shop{}

	errGet := db.First(&shop, id).Error
	if errGet != nil {
		return shop, errGet
	}

	return shop, nil
}
