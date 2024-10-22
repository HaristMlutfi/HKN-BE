package services

import (
	"hkn-be/config"
	"hkn-be/infras/jwt_infra"
	"hkn-be/infras/mail"
	"hkn-be/repositories"
	"hkn-be/services/auth_service"
	"hkn-be/services/user_service"

	"github.com/redis/go-redis/v9"
)

type ServiceCtx struct {
	AuthService auth_service.AuthServiceInterface
	UserService user_service.UserService
}

func InitServices(
	ctxRepo *repositories.RepositoryCtx,
	server config.ServerConfig,
	jwtInterface jwt_infra.JwtInterface,
	redis *redis.Client,
	mail mail.MailServiceInterface,

) *ServiceCtx {
	authService := auth_service.NewAuthService(ctxRepo, server, jwtInterface, redis, mail)
	userService := user_service.NewUserService(ctxRepo.UserRepo)
	return &ServiceCtx{
		AuthService: authService,
		UserService: userService,
	}
}
