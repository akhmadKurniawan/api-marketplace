package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type ProductTypeRepository struct {
	DB         *gorm.DB
	Collection *mongo.Collection
}

func NewProductTypeRepository(db *gorm.DB, database *mongo.Database) infrastructure.ProductTypeRepository {
	return &ProductTypeRepository{
		DB:         db,
		Collection: database.Collection("product_types"),
	}
}

func (repo *ProductTypeRepository) CreateProductType(ctx context.Context, productType models.ProductType) error {
	db := repo.DB
	dbm := repo.Collection

	_, err := dbm.InsertOne(ctx, productType)
	if err != nil {
		return err
	}

	errCreate := db.Create(&productType).Error
	if errCreate != nil {
		return errCreate
	}

	return nil
}

func (repo *ProductTypeRepository) GetProductType(ctx context.Context, params models.ProductType) (models.ProductType, error) {
	db := repo.DB
	productType := models.ProductType{}

	if params.Name != "" {
		db = db.Where("name = ?", params.Name)
	}

	errGet := db.Find(&productType).Error
	if errGet != nil {
		return productType, nil
	}

	return productType, nil
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
