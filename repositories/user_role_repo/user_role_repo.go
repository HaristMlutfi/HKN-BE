package user_role_repo

import (
	"context"
	"hkn-be/constants"
	"hkn-be/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userRoleRepo struct {
	*gorm.DB
}

func (u userRoleRepo) GetUserRoleListByUserId(ctx context.Context, userId string) (*[]models.UserRole, error) {
	var result []models.UserRole
	tx := u.DB.WithContext(ctx).Table(models.TableNameUserRole).Find(&result, constants.QueryUserId, userId)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, tx.Error
	}
	return &result, nil
}

func (u userRoleRepo) UpdateUserRole(ctx context.Context, userRole models.UserRole) error {
	tx := u.DB.WithContext(ctx).Table(models.TableNameUserRole).Updates(&userRole)
	if tx.Error != nil {
		log.Error(tx.Error)
		return tx.Error
	}
	return nil
}

func (u userRoleRepo) CreateUserRole(ctx context.Context, userRole models.UserRole) (*models.UserRole, error) {
	tx := u.WithContext(ctx).Create(&userRole)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, tx.Error
	}
	return &userRole, nil
}

func (u userRoleRepo) DeleteUserRole(ctx context.Context, userRole models.UserRole) error {
	tx := u.DB.WithContext(ctx).Table(models.TableNameUserRole).Delete(&userRole)
	if tx.Error != nil {
		log.Error(tx.Error)
		return tx.Error
	}
	return nil
}
