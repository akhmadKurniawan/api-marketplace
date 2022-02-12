package repository

import (
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

func (p *loginRepository) Login(ctx context.Context, accessToken models.UserToken, userID uint64) (models.UserToken, error) {
	accToken := models.UserToken{}
	user := models.User{}
	userData := models.User{}
	user.LastLogin = time.Now()

	errUpUser := p.DB.Model(&userData).Where("id = ?", userID).Update(&user).Error
	if errUpUser != nil {
		return accessToken, errUpUser
	}

	checkToken := p.DB.Where("user_id = ?", userID).First(&accToken).RecordNotFound()
	if checkToken {
		err := p.DB.Create(&accessToken).Error
		if err != nil {
			return accessToken, err
		}
		return accessToken, nil

	} else {
		if err := p.DB.Model(&accToken).Where("user_id = ?", userID).Update(&accessToken).Error; err != nil {
			return accessToken, err
		}
		return accessToken, nil
	}
}
