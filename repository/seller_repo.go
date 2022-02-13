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

func (repo *SellerRepository) CreateSeller(ctx context.Context, model models.Seller) error {
	return nil
}
