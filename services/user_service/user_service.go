package user_service

import (
	"context"
	"hkn-be/repositories/user_repo"
)

type userService struct {
	userRepo user_repo.UserRepoInterface
}

func NewUserService(userRepo user_repo.UserRepoInterface) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {

	return s.userRepo.DeleteUser(ctx, id)
}
