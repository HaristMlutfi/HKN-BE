package repositories

import (
	"gorm.io/gorm"
	"hkn-be/repositories/role_repo"
	"hkn-be/repositories/user_repo"
	"hkn-be/repositories/user_role_repo"
)

type RepositoryCtx struct {
	DB           *gorm.DB
	UserRepo     user_repo.UserRepoInterface
	UserRoleRepo user_role_repo.UserRoleRepoInterface
	RoleRepo     role_repo.RoleRepoInterface
}

func InitRepositories(db *gorm.DB) *RepositoryCtx {
	userRepo := user_repo.NewUserRepo(db)
	userRoleRepo := user_role_repo.NewUserRoleRepo(db)
	roleRepo := role_repo.NewRoleRepo(db)
	return &RepositoryCtx{
		DB:           db,
		UserRepo:     userRepo,
		UserRoleRepo: userRoleRepo,
		RoleRepo:     roleRepo,
	}
}
