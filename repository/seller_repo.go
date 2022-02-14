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
	user := models.User{}
	seller := models.Seller{}

	tx := repo.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&seller).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("user_id = ?", id).Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
