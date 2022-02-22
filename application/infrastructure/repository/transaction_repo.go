package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) infrastructure.TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (repo *TransactionRepository) CreateTransaction(ctx context.Context, transaction models.Transaction) error {
	db := repo.DB

	tx := db.Begin()

	errCreate := tx.Create(&transaction).Error
	if errCreate != nil {
		tx.Rollback()
		return errCreate
	}

	tx.Commit()
	return nil

}
