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

func (repo *TransactionRepository) UpdateScheduler(ctx context.Context, param models.Transaction) error {
	db := repo.DB.Debug()
	transactionData := models.Transaction{}

	errUpdate := db.Raw("UPDATE transactions SET status = ?, updated_at = ? WHERE status = 'pending' AND created_at + interval 1 hour < now()", param.Status, param.UpdatedAt).Scan(transactionData).Error
	// errUpdate := db.Model(&transactionData).Where("status = pending AND created_at + interval 1 hour < now()").Updates(models.Transaction{Status: param.Status, UpdatedAt: param.UpdatedAt}).Error
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (repo *TransactionRepository) UpdateTransaction(ctx context.Context, transaction models.Transaction, id string) (models.Transaction, error) {
	db := repo.DB
	transactionData := models.Transaction{}

	errUpdate := db.Model(&transactionData).Where("id_va = ?", id).Update("status", transaction.Status).Error
	if errUpdate != nil {
		return transaction, errUpdate
	}

	return transaction, nil
}
