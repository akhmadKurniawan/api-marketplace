package repository

import (
	"app/application/infrastructure"
	"app/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type loginRepository struct {
	DB *gorm.DB
}

func NewLoginRepository(DB *gorm.DB) infrastructure.LoginRepository {
	return &loginRepository{
		DB: DB,
	}
}

func (repo *loginRepository) Login(ctx context.Context, accessToken models.UserToken, userID int) (models.UserToken, error) {
	db := repo.DB.Debug()
	accToken := models.UserToken{}
	user := models.User{}
	user.LastLoginAt = time.Now()

	errUpUser := db.Model(&user).Where("id = ?", userID).Updates(&user).Error
	if errUpUser != nil {
		return accessToken, errUpUser
	}

	checkToken := db.Where("user_id = ?", userID).First(&accToken).Error
	if checkToken != nil {
		err := db.Create(&accessToken).Error
		if err != nil {
			return accessToken, err
		}
		return accessToken, nil
	} else {
		if err := db.Model(&accToken).Where("user_id = ?", userID).Updates(&accessToken).Error; err != nil {
			return accessToken, err
		}
		return accessToken, nil
	}
}
