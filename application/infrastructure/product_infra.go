package infrastructure

import (
	"app/models"
	"context"
)

type ProductRepository interface {
	CreateProduct(context.Context, models.Product) error
}
