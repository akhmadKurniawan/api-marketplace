package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"

	"gorm.io/gorm"
)

type CostumerRepository struct {
	DB *gorm.DB
}

func NewCostumerRepository(db *gorm.DB) infrastructure.CostumerRepository {
	return &CostumerRepository{
		DB: db,
	}
}

func (repo *CostumerRepository) CreateCostumer(ctx context.Context, costumer models.Costumer) error {
	db := repo.DB

	errCreate := db.Create(&costumer).Error
	if errCreate != nil {
		return errCreate
	}

	return nil
}
