package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB  *gorm.DB
	DBM *mongo.Database
}

func NewProductRepository(db *gorm.DB, dbm *mongo.Database) infrastructure.ProductRepository {
	return &ProductRepository{
		DB:  db,
		DBM: dbm,
	}
}

func (repo *ProductRepository) CreateProduct(ctx context.Context, product models.Product) error {
	db := repo.DB
	dbm := repo.DBM.Collection("products")

	go func() {
		_, err := dbm.InsertOne(ctx, product)
		if err != nil {
			log.Println(err)
		}
	}()

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

	errGet := db.Find(&product).Error
	if errGet != nil {
		return product, errGet
	}

	return product, nil
}
