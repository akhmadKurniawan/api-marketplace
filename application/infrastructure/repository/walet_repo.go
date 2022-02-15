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
