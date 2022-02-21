package infrastructure

import (
	"app/models"
	"context"
)

type WaletRepository interface {
	CreateWalet(context.Context, models.Walet) error
	GetWaletByUserID(context.Context, int) (models.Walet, error)
	UpdateWaletSaldo(context.Context, int, int) (models.Walet, error)
}
