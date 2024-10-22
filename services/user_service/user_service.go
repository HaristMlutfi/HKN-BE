package user_service

import (
	"context"
	"hkn-be/models"
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

func (s *userService) UpdateUser(ctx context.Context, user models.User) error {

	return s.userRepo.UpdateUser(ctx, user)
}
