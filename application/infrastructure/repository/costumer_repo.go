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
	db := repo.DB.Debug()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(&costumer).Error; err != nil {
		tx.Rollback()
		return err
	}

	// userID := db.Where("user_id = ?", costumer.UserID).Take(&costumer).Error
	// if userID != nil {
	// 	db.Create(&costumer)
	// }
	// return nil
	return tx.Commit().Error
}

func (repo *CostumerRepository) GetCostumerByUserId(ctx context.Context, id int) (models.Costumer, error) {
	costumer := models.Costumer{}
	db := repo.DB

	if err := db.Where("user_id = ?", id).First(&costumer).Error; err != nil {
		return costumer, err
	}
	return costumer, nil
}
