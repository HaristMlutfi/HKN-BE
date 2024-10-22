package user_service

import (
	"context"
	"hkn-be/models"
)

type UserService interface {
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, user models.User) error
}
