package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"errors"
	"log"

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
	if err := tx.Create(&seller).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (repo *SellerRepository) GetSellerByID(ctx context.Context, id int) (models.Seller, error) {
	db := repo.DB
	seller := models.Seller{}

	errGet := db.First(&seller, id).Error
	if errGet != nil {
		return seller, errGet
	}

	return seller, nil
}

func (repo *SellerRepository) GetSellerByUserID(ctx context.Context, id int) (models.Seller, error) {
	seller := models.Seller{}

	if err := repo.DB.Where("user_id = ?", id).First(&seller).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Repository - GetSellerByUserID Error : ", err)
		}
		return seller, err
	}

	return seller, nil
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
