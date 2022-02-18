package infrastructure

import (
	"app/models"
	"context"
)

type ProductRepository interface {
	CreateProduct(context.Context, models.Product, string) error
	GetProduct(context.Context, models.Product) (models.Product, error)
}
