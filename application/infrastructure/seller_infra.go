package infrastructure

import (
	"app/models"
	"context"
)

type SellerRepository interface {
	CreateSeller(context.Context, models.Seller) error
}
