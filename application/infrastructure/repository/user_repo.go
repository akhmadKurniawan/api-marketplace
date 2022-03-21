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

func (repo *UserRepository) SignUpUser(ctx context.Context, user models.User) (models.User, error) {
	db := repo.DB.Debug()
	user.LastLoginAt = time.Now()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return models.User{}, err
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	return user, tx.Commit().Error

}

func (repo *UserRepository) GetUserID(ctx context.Context, id string) (models.User, error) {
	db := repo.DB.Debug()

	user := models.User{}

	if err := db.Preload("UserToken").First(&user, id).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Repository - GetUserID Error : ", err)
		}
		return user, err
	}
	return user, nil
}

func (repo *UserRepository) GetUsername(ctx context.Context, username string) (models.User, error) {
	db := repo.DB.Debug()
	userData := models.User{}

	if err := db.Where("username = ?", username).First(&userData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Repository - GetUsername Error : ", err)
		}
		return userData, err
	}
	return userData, nil
}

func (repo *UserRepository) GetStatus(ctx context.Context, status string) (models.User, error) {
	db := repo.DB.Debug()
	userData := models.User{}

	if err := db.Where("status = ?", status).Find(&userData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Repository - GetStatus error :", err)
		}
		return userData, err
	}
	return userData, nil
}

func (repo *UserRepository) GetAllUsername(ctx context.Context, username string) (models.User, error) {
	db := repo.DB.Debug()
	userData := models.User{}

	if err := db.Where("username = ?", username).Find(&userData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Repository - GetUsername Error : ", err)
		}
		return userData, err
	}
	return userData, nil
}

func (repo *UserRepository) GetUserTransaction(ctx context.Context, param models.User) (*models.User, error) {
	db := repo.DB
	userData := models.User{}

	errGet := db.First(&userData, "id = ?", param.ID).Error
	if errGet != nil {
		if !errors.Is(errGet, gorm.ErrRecordNotFound) {
			log.Println("UserRepository - GetUser Error :", errGet)
		}
		return nil, errGet
	}
	return &userData, nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, user models.User, id string) (models.User, error) {
	userData := models.User{}
	errUpdate := repo.DB.Model(&userData).Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		return user, errUpdate
	}
	return user, nil
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

func (repo *UserRepository) VerifyEmailUser(ctx context.Context, id string, param models.User) error {
	db := repo.DB
	user := models.User{}

	errUpdate := db.Model(&user).Where("id = ?", id).Update("status", param.Status).Error
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
