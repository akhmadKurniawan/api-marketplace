package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"

	"gorm.io/gorm"
)

type SellerRepository struct {
	DB *gorm.DB
}

func NewSellerRepository(db *gorm.DB) infrastructure.SellerRepository {
	return &SellerRepository{
		DB: db,
	}
}

func (repo *SellerRepository) CreateSeller(ctx context.Context, seller models.Seller) error {
	db := repo.DB

	errCreate := db.Create(&seller).Error
	if errCreate != nil {
		return errCreate
	}
	return nil
}

func (repo *SellerRepository) DeleteSeller(ctx context.Context, id string) error {
	seller := models.Seller{}
	db := repo.DB

	errDelete := db.Where("id = ?", id).Delete(&seller).Error
	if errDelete != nil {
		return errDelete
	}
	return nil
}
