package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) infrastructure.ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (repo *ProductRepository) CreateProduct(ctx context.Context, product models.Product, imageFile string) error {
	db := repo.DB
	shop := models.Shop{}
	productType := models.ProductType{}

	errShop := db.First(&shop, product.ShopId).Error
	if errShop != nil {
		return errors.New("shop tidak ditemukan")
	}

	errShopN := db.Where("name = ? AND shop_id = ?", product.Name, product.ShopId).Find(&product).Error
	if errShopN != nil {
		return errors.New("name already exist")
	}

	errProductType := db.First(&productType, product.ProductType).Error
	if errProductType != nil {
		return errors.New("product type tidak ditemukan")
	}

	product.Image = imageFile
	product.ProductType = productType.ID
	product.ShopId = shop.ID

	errCreate := db.Create(&product).Error
	if errCreate != nil {
		return errCreate
	}

	return nil
}
