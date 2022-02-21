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

func (repo *WaletRepository) UpdateWaletSaldo(ctx context.Context, id int, saldo int) (models.Walet, error) {
	db := repo.DB
	walet := models.Walet{}

	tx := db.Begin()
	db.Debug().Model(&walet).Where("user_id = ?", id).Update("saldo", saldo)
	errUp := tx.Model(&walet).Where("user_id = ?", id).Update("saldo", saldo).Error
	if errUp != nil {
		tx.Rollback()
		return walet, errUp
	}

	tx.Commit()
	return walet, nil
}

// func (repo *WaletRepository) UpdateWaletSaldoSeller(ctx context.Context, id int, saldo int) (models.Walet,error) {
// 	db := repo.DB
// 	walet := models.Walet{}
// 	seller := models.Seller{}

// 	tx := db.Begin()
// 	errUp := tx.Model(&walet).Where("user_id = ?", id).Model(&seller).Where("user_id = ?", id)
// }
