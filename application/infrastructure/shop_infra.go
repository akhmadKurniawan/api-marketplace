package infrastructure

import (
	"app/models"
	"context"
)

type ShopRepository interface {
	CreateShop(context.Context, models.Shop) error
	GetShopById(context.Context, int) (models.Shop, error)
}
