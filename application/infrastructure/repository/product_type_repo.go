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

func (repo *ProductTypeRepository) GetProductTypeById(ctx context.Context, id int) (models.ProductType, error) {
	db := repo.DB
	productType := models.ProductType{}

	errGet := db.First(&productType, id).Error
	if errGet != nil {
		return productType, errGet
	}

	return productType, nil
}
