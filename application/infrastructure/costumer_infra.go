package infrastructure

import (
	"app/models"
	"context"
)

type CostumerRepository interface {
	CreateCostumer(context.Context, models.Costumer, int) error
}
