package infrastructure

import (
	"app/models"
	"context"
)

type TransactionRepository interface {
	CreateTransaction(context.Context, models.Transaction) error
}
