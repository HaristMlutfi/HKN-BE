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

// file: services/user_service.go
func (s *userService) UpdateUser(ctx context.Context, userId string, user *models.User) error {
	// Cari data user yang ada
	existingUser, err := s.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	// Periksa apakah password baru disediakan
	if user.Password != "" {
		// Hash password baru jika ada
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return err // Kembalikan error jika hashing gagal
		}
		// Update password yang sudah di-hash
		existingUser.Password = hashedPassword
	}

	// Update field lain seperti nama
	existingUser.Name = user.Name

	// Simpan data yang sudah diperbarui ke database
	return s.userRepo.UpdateUser(ctx, *existingUser)
}

// userService.go
func (s *userService) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	return s.userRepo.GetUserById(ctx, userId)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(ctx, email)
}
