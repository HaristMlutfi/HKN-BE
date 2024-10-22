package user_service

import "context"

type UserService interface {
	DeleteUser(ctx context.Context, id string) error
}
