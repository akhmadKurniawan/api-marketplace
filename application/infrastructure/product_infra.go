package infrastructure

import (
	"app/models"
	"context"
)

type ProductRepository interface {
	CreateProduct(context.Context, models.Product) error
	GetProductByID(context.Context, int) (models.Product, error)
	GetProduct(context.Context, models.Product) (models.Product, error)
}
