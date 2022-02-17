package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"

	"gorm.io/gorm"
)

type ProductTypeRepository struct {
	DB *gorm.DB
}

func NewProductTypeRepository(db *gorm.DB) infrastructure.ProductTypeRepository {
	return &ProductTypeRepository{
		DB: db,
	}
}

func (repo *ProductTypeRepository) CreateProductType(ctx context.Context, productType models.ProductType, img string) error {
	db := repo.DB

	productType.Image = img
	errCreate := db.Create(&productType).Error
	if errCreate != nil {
		return errCreate
	}

	return nil
}
