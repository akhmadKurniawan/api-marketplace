package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) infrastructure.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repo *UserRepository) SignUpUser(ctx context.Context, user models.User) error {
	db := repo.DB
	user.LastLoginAt = time.Now()

	errCreate := db.Create(&user).Error
	if errCreate != nil {
		return errCreate
	}
	return nil
}
