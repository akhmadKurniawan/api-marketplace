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

func (repo *ProductTypeRepository) CreateProductType(ctx context.Context, productType models.ProductType) error {
	db := repo.DB

	errCreate := db.Create(&productType).Error
	if errCreate != nil {
		return errCreate
	}

	return nil
}
