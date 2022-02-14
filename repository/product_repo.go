package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"

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

func (repo *ProductRepository) CreateProduct(ctx context.Context, product models.Product) error {
	db := repo.DB

	errCreate := db.Create(&product).Error
	if errCreate != nil {
		return errCreate
	}

	return nil
}
