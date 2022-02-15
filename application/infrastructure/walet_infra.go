package infrastructure

import (
	"app/models"
	"context"
)

type WaletRepository interface {
	CreateWalet(context.Context, models.Walet) error
}
