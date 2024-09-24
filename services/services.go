package services

import (
	"github.com/redis/go-redis/v9"
	"hkn-be/config"
	"hkn-be/infras/jwt_infra"
	"hkn-be/infras/mail"
	"hkn-be/repositories"
	"hkn-be/services/auth_service"
)

type ServiceCtx struct {
	AuthService auth_service.AuthServiceInterface
}

func InitServices(
	ctxRepo *repositories.RepositoryCtx,
	server config.ServerConfig,
	jwtInterface jwt_infra.JwtInterface,
	redis *redis.Client,
	mail mail.MailServiceInterface,

) *ServiceCtx {
	authService := auth_service.NewAuthService(ctxRepo, server, jwtInterface, redis, mail)
	return &ServiceCtx{
		AuthService: authService,
	}
}
