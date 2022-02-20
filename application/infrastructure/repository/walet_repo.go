package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"

	"gorm.io/gorm"
)

type WaletRepository struct {
	DB *gorm.DB
}

func NewWaletRepository(db *gorm.DB) infrastructure.WaletRepository {
	return &WaletRepository{
		DB: db,
	}
}

func (repo *WaletRepository) CreateWalet(ctx context.Context, walet models.Walet) error {
	db := repo.DB

	errCreate := db.Create(&walet).Error
	if errCreate != nil {
		return errCreate
	}
	return nil
}

func (repo *WaletRepository) GetWaletByUserID(ctx context.Context, id int) (models.Walet, error) {
	db := repo.DB
	walet := models.Walet{}

	errGet := db.Where("user_id = ?", id).First(&walet).Error
	if errGet != nil {
		return walet, errGet
	}

	return walet, nil
}

func (repo *WaletRepository) UpdateWaletByUserID(ctx context.Context, id int) (models.Walet, error) {
	db := repo.DB
	walet := models.Walet{}

	tx := db.Begin()
	errUp := tx.Where("user_id = ?", id).Update("saldo", walet.Saldo).Error
	if errUp != nil {
		tx.Rollback()
	}

	return walet, nil
}
