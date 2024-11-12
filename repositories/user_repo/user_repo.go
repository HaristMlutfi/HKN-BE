package user_repo

import (
	"context"

	"hkn-be/constants"
	"hkn-be/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userRepo struct {
	*gorm.DB
}

func (u userRepo) UpdateUser(ctx context.Context, user models.User) error {
	tx := u.DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", user.Id).Updates(&user)
	if tx.Error != nil {
		log.Error("Failed to update user:", tx.Error)
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

func (u userRepo) DeleteUser(ctx context.Context, userId string) error {
	tx := u.DB.WithContext(ctx).Unscoped().Delete(&models.User{}, constants.QueryById, userId)
	if tx.Error != nil {
		log.Error(tx.Error)
		return tx.Error
	}
	return nil
}

// userRepo.go
func (u userRepo) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	var user models.User
	err := u.DB.WithContext(ctx).Where("id = ?", userId).First(&user).Error
	if err != nil {
		log.Error("User not found:", err)
		return nil, err
	}
	return &user, nil
}

func (u userRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Error("User not found:", err)
		return nil, err
	}
	return &user, nil
}
