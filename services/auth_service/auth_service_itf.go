package auth_service

import (
	"context"
	"github.com/redis/go-redis/v9"
	"hkn-be/config"
	"hkn-be/infras/jwt_infra"
	"hkn-be/infras/mail"
	"hkn-be/objects"
	"hkn-be/repositories"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, userData objects.User) error
	RequestVerificationCode(ctx context.Context, email string) error
	VerifyUserRegistration(ctx context.Context, request objects.VerifyRegistrationRequest) error
	Login(ctx context.Context, loginRequestData objects.LoginRequest) (*objects.LoginResponse, error)
	Logout(ctx context.Context, id string) error
}

func NewAuthService(
	repoCtx *repositories.RepositoryCtx, server config.ServerConfig, jwtInterface jwt_infra.JwtInterface, redis *redis.Client,
	mailService mail.MailServiceInterface,
) AuthServiceInterface {
	return &authService{
		repoCtx, server, jwtInterface, redis, mailService,
	}
}
