package user_repo

import (
	"context"
	"hkn-be/models"

	"gorm.io/gorm"
)

type UserRepoInterface interface {
	GetUserById(ctx context.Context, userId string) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	DeleteUser(ctx context.Context, userId string) error
}

func NewUserRepo(db *gorm.DB) UserRepoInterface {
	return &userRepo{
		db,
	}
}
