package infrastructure

import (
	"app/models"
	"context"
)

type ShopRepository interface {
	CreateShop(context.Context, models.Shop) error
	GetShopBySellerID(ctx context.Context, id int) (models.Shop, error)
}
