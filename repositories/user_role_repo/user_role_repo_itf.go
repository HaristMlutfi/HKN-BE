package user_role_repo

import (
	"context"
	"gorm.io/gorm"
	"hkn-be/models"
)

type UserRoleRepoInterface interface {
	GetUserRoleListByUserId(ctx context.Context, userId string) (*[]models.UserRole, error)
	UpdateUserRole(ctx context.Context, user models.UserRole) error
	CreateUserRole(ctx context.Context, user models.UserRole) (*models.UserRole, error)
	DeleteUserRole(ctx context.Context, user models.UserRole) error
}

func NewUserRoleRepo(db *gorm.DB) UserRoleRepoInterface {
	return &userRoleRepo{
		db,
	}
}
