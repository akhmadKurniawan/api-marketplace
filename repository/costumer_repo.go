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

func (repo *CostumerRepository) CreateCostumer(ctx context.Context, costumer models.Costumer, userID int) error {
	db := repo.DB

	userId := db.Where("user_id = ?", userID).First(&costumer).Error
	if userId != nil {
		errCreate := db.Create(&costumer).Error
		if errCreate != nil {
			return errCreate
		}
	}

	return nil
}
