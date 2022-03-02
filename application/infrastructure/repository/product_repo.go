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

func (repo *ProductRepository) GetProductByShopID(ctx context.Context, id int) ([]models.Product, error) {
	db := repo.DB
	product := []models.Product{}

	errGet := db.Where("shop_id = ?", id).Find(&product).Error
	if errGet != nil {
		return nil, errGet
	}

	return product, nil
}

func (repo *ProductRepository) GetProductByID(ctx context.Context, id int) (models.Product, error) {
	db := repo.DB
	product := models.Product{}

	errGet := db.First(&product, id).Error
	if errGet != nil {
		return product, errGet
	}

	return product, nil
}

func (repo *ProductRepository) UpdateProdut(ctx context.Context, id int, qty int) (models.Product, error) {
	db := repo.DB.Debug()
	product := models.Product{}

	tx := db.Begin()

	errUp := tx.Model(&product).Where("id = ?", id).Update("qty", qty).Error
	if errUp != nil {
		tx.Rollback()
		return product, errUp
	}

	tx.Commit()
	return product, nil
}

func (repo *ProductRepository) GetProduct(ctx context.Context, params models.Product) (models.Product, error) {
	db := repo.DB
	product := models.Product{}

	if params.Name != "" {
		db = db.Where("name = ?", params.Name)
	}
	if params.ShopId != 0 {
		db = db.Where("shop_id = ?", params.ShopId)
	}
	if params.Qty != 0 {
		db = db.Where("qty = ?", params.Qty)
	}
	if params.Price != 0 {
		db = db.Where("price = ?", params.Price)
	}

	db.Debug().Find(&product)
	errGet := db.Find(&product).Error
	if errGet != nil {
		return product, errGet
	}

	return product, nil
}
