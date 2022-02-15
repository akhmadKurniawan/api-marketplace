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

	errCreate := db.Create(&transaction).Error
	if errCreate != nil {
		return errCreate
	}

	return nil
}
