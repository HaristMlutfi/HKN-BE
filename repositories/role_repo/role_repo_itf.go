package role_repo

import (
	"context"
	"gorm.io/gorm"
	"hkn-be/models"
)

type RoleRepoInterface interface {
	GetRoleById(ctx context.Context, roleId string) (*models.Role, error)
	GetRoleByCode(ctx context.Context, code string) (*models.Role, error)
}

func NewRoleRepo(db *gorm.DB) RoleRepoInterface {
	return &roleRepo{
		db,
	}
}
