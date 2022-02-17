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

func (repo *UserRepository) DeleteUser(ctx context.Context, id string) error {
	user := models.User{}
	seller := models.Seller{}
	costumer := models.Costumer{}
	token := models.UserToken{}

	tx := repo.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("user_id = ?", id).Delete(&seller).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("user_id = ?", id).Delete(&costumer).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("user_id = ?", id).Delete(&token).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
