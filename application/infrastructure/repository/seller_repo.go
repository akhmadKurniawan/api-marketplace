package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"errors"
	"fmt"
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
	db := repo.DB

	errCreate := db.Create(&seller).Error
	if errCreate != nil {
		return errCreate
	}
	return nil
}

func (repo *SellerRepository) GetSellerByUserID(ctx context.Context, id int) (models.Seller, error) {
	seller := models.Seller{}

	if err := repo.DB.Preload("Shop").First(&seller, id).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Repository - GetSellerByUserID Error : ", err)
		}
		return seller, err
	}

	fmt.Println(seller)
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
