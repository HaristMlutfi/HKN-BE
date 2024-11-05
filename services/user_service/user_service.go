package user_service

import (
	"context"
	"hkn-be/models"
	"hkn-be/repositories/user_repo"
	"hkn-be/utils"
)

type userService struct {
	userRepo user_repo.UserRepoInterface
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.DeleteUser(ctx, id)
}

func (s *userService) CreateUser(ctx context.Context, user models.User) error {
	// Gunakan fungsi HashPassword dari utils
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err // Kembalikan error jika hashing gagal
	}

	// Simpan hashed password ke struct user
	user.Password = hashedPassword

	// Panggil repository untuk menyimpan user
	_, err = s.userRepo.CreateUser(ctx, user)
	return err
}

func (s *userService) UpdateUser(ctx context.Context, user models.User) error {
	return s.userRepo.UpdateUser(ctx, user)
}

func (s *userService) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	return s.userRepo.GetUserById(ctx, userId)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(ctx, email)
}
