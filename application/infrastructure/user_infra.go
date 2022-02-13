package infrastructure

import (
	"app/models"
	"context"
)

type UserRepository interface {
	SignUpUser(context.Context, models.User) error
	GetUserID(ctx context.Context, id string) (models.User, error)
	GetUsername(context.Context, string) (models.User, error)
}
