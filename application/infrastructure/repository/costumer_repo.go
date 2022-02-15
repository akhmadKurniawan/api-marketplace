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

	userID := db.Where("user_id = ?", costumer.UserID).Take(&costumer).Error
	if userID != nil {
		db.Create(&costumer)
	}

	return nil
}
