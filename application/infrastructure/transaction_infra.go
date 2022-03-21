package infrastructure

import (
	"app/models"
	"context"
)

type TransactionRepository interface {
	CreateTransaction(context.Context, models.Transaction) error
	GetTransactions(context.Context) ([]models.Transaction, error)
	UpdateTransaction(context.Context, models.Transaction, string) (models.Transaction, error)
	UpdateScheduler(context.Context, models.Transaction) error
}
