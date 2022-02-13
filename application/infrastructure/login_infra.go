package infrastructure

import (
	"app/models"
	"context"
)

type LoginRepository interface {
	Login(context.Context, models.UserToken, int) (models.UserToken, error)
}
