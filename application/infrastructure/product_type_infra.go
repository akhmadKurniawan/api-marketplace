package infrastructure

import (
	"app/models"
	"context"
)

type ProductTypeRepository interface {
	CreateProductType(context.Context, models.ProductType) error
	GetProductType(context.Context, models.ProductType) (models.ProductType, error)
	GetProductTypeById(context.Context, int) (models.ProductType, error)
}
