package infrastructure

import (
	"app/models"
	"context"
)

type UserRepository interface {
	DeleteUser(context.Context, string) error
	SignUpUser(context.Context, models.User) (models.User, error)
	GetUsername(context.Context, string) (models.User, error)
	GetAllUsername(context.Context, string) (models.User, error)
	GetUserID(ctx context.Context, id string) (models.User, error)
	UpdateUser(context.Context, models.User, string) (models.User, error)
	VerifyEmailUser(context.Context, string, models.User) error
}
