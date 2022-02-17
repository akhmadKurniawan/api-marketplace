package infrastructure

import (
	"app/models"
	"context"
)

type ProductTypeRepository interface {
	CreateProductType(context.Context, models.ProductType, string) error
}
