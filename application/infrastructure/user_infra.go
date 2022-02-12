package infrastructure

import (
	"app/models"
	"context"
)

type UserRepository interface {
	SignUpUser(ctx context.Context, user models.User) error
}
