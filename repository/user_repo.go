package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"errors"
	"log"
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

func (repo *UserRepository) GetUserID(ctx context.Context, id string) (models.User, error) {
	user := models.User{}

	if err := repo.DB.Preload("UserToken").First(&user, id).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Repository - GetUserID Error : ", err)
		}
		return user, err
	}
	return user, nil
}

func (repo *UserRepository) GetUsername(ctx context.Context, username string) (models.User, error) {
	userData := models.User{}

	if err := repo.DB.Where("username = ?", username).First(&userData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Repository - GetUsername Error : ", err)
		}
		return userData, err
	}
	return userData, nil
}
