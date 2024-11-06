package user_service

import (
	"context"
	"hkn-be/models"
	"hkn-be/repositories/user_repo"
)

type UserServiceInterface interface {
	DeleteUser(ctx context.Context, id string) error
	CreateUser(ctx context.Context, user models.User) error
	UpdateUser(ctx context.Context, id string, user *models.User) error
	GetUserById(ctx context.Context, userId string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

func NewUserService(userRepo user_repo.UserRepoInterface) UserServiceInterface {
	return &userService{
		userRepo: userRepo,
	}
}
