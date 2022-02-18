package infrastructure

import (
	"app/models"
	"context"
)

type ShopRepository interface {
	CreateShop(context.Context, models.Shop, string) error
	GetShopById(context.Context, int) (models.Shop, error)
}
