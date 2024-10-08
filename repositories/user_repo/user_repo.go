package user_repo

import (
	"context"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"hkn-be/constants"
	"hkn-be/models"
)

type userRepo struct {
	*gorm.DB
}

func (u userRepo) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	var result models.User
	tx := u.DB.WithContext(ctx).Table(models.TableNameUser).First(&result, constants.QueryById, userId)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, tx.Error
	}
	return &result, nil
}

func (u userRepo) UpdateUser(ctx context.Context, user models.User) error {
	tx := u.DB.WithContext(ctx).Table(models.TableNameUser).Updates(&user)
	if tx.Error != nil {
		log.Error(tx.Error)
		return tx.Error
	}
	return nil
}

func (u userRepo) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	tx := u.WithContext(ctx).Create(&user)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, tx.Error
	}
	return &user, nil
}

func (u userRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var result models.User
	tx := u.DB.WithContext(ctx).First(&result, constants.QueryEmail, email)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, tx.Error
	}
	return &result, nil
}
