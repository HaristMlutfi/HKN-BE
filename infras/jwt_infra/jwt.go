package jwt_infra

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"hkn-be/config"
	"hkn-be/constants"
	"hkn-be/objects"
	"net/http"
	"time"
)

const (
	AuthKey = "LMS-app-2024"
	bearer  = "Bearer"
)

type jwtObj struct {
	config *config.JwtConfig
	redis  *redis.Client
}

type JwtCustomClaims struct {
	jwt.RegisteredClaims
	Id    string
	Email string
	Name  string
	Roles []objects.Role
}

type JwtRequest struct {
	ID    string
	Email string
	Name  string
}

type JwtInterface interface {
	GenerateJwtToken(ctx context.Context, data objects.User) (*string, error)
	//ExtractJwtClaims(ctx context.Context, authBearer string) (*JwtCustomClaims, error)
	ValidateTokenIssuer(claims *JwtCustomClaims) error
	ValidateTokenExpire(ctx context.Context, claims *JwtCustomClaims, reqToken string) error

	GetEchoJwtMiddlewareConfig() echo.MiddlewareFunc
	RoleBasedAuth(roles ...string) echo.MiddlewareFunc

	SaveTokenToRedis(ctx context.Context, hour int, id, token, authKey string) error
	GetTokenFromRedis(ctx context.Context, id, authKey string) (*string, error)
	DeleteTokenFromRedis(ctx context.Context, id, authKey string) error
}

func NewJwt(cfg *config.JwtConfig, redis *redis.Client) JwtInterface {
	return &jwtObj{
		config: cfg,
		redis:  redis,
	}
}

func (j jwtObj) GetEchoJwtMiddlewareConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			ContextKey: "authKey",
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(JwtCustomClaims)
			},
			SigningKey: []byte(j.config.Secret),
		},
	)
}

func (j jwtObj) RoleBasedAuth(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get("authKey").(*jwt.Token).Claims.(*JwtCustomClaims)
			userRoles := claims.Roles

			for _, role := range roles {
				for _, userRole := range userRoles {
					if role == userRole.Code {
						return next(c)
					}
				}
			}
			return echo.NewHTTPError(http.StatusForbidden, "You don't have the right permissions")
		}
	}
}

func (j jwtObj) GenerateJwtToken(ctx context.Context, data objects.User) (*string, error) {
	expireTime := time.Now().Add(time.Duration(j.config.TokenLifeTimeHour) * time.Hour)

	claims := &JwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.config.Issuer,
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Id:    data.Id,
		Email: data.Email,
		Name:  data.Name,
		Roles: data.Roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.config.Secret))
	if err != nil {

		return nil, err
	}
	err = j.SaveTokenToRedis(ctx, j.config.TokenLifeTimeHour, data.Id, signedToken, AuthKey)
	if err != nil {
		log.Error(constants.ErrSaveTokenToRedis)
		return nil, constants.ErrSaveTokenToRedis
	}
	return &signedToken, nil
}

//func (j jwtObj) ExtractJwtClaims(ctx context.Context, authBearer string) (*JwtCustomClaims, error) {
//	splitToken := strings.Split(authBearer, bearer)
//	if len(splitToken) != 2 {
//		return nil, consterr.ErrTokenRequired
//	}
//	reqToken := strings.TrimSpace(splitToken[1])
//	t, err := jwt.ParseWithClaims(
//		reqToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//			return j.config.Secret, nil
//		},
//	)
//
//	if err != nil && err.Error() != consterr.ErrKeyIsNotInvalidType.Error() {
//		return nil, err
//	}
//	claims := t.Claims.(*JwtCustomClaims)
//
//	err = j.ValidateTokenIssuer(claims)
//	if err != nil {
//		return nil, err
//	}
//	err = j.ValidateTokenExpire(ctx, claims, reqToken)
//	if err != nil {
//		return nil, err
//	}
//	return claims, nil
//}

func (j jwtObj) ValidateTokenIssuer(claims *JwtCustomClaims) error {
	if claims.Issuer != j.config.Issuer {
		return constants.ErrInvalidIssuer
	}
	return nil
}

func (j jwtObj) ValidateTokenExpire(ctx context.Context, claims *JwtCustomClaims, reqToken string) error {
	token, err := j.GetTokenFromRedis(ctx, claims.ID, AuthKey)
	if err != nil {
		log.Error(constants.ErrGetTokenFromRedis)
		return constants.ErrGetTokenFromRedis
	}

	if *token == "" {
		return constants.ErrTokenAlreadyExpired
	}
	if *token != reqToken {
		return constants.ErrTokenAlreadyExpired
	}
	return nil
}

func (j jwtObj) SaveTokenToRedis(ctx context.Context, hour int, id, token, authKey string) error {
	key := fmt.Sprintf("%s:%d", authKey, id)
	ttl := time.Duration(hour) * time.Hour
	err := j.redis.Set(ctx, key, token, ttl).Err()
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (j jwtObj) GetTokenFromRedis(ctx context.Context, id, authKey string) (*string, error) {
	key := fmt.Sprintf("%s:%d", authKey, id)
	val, err := j.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return &val, nil
}

func (j jwtObj) DeleteTokenFromRedis(ctx context.Context, id, authKey string) error {
	key := fmt.Sprintf("%s:%d", authKey, id)
	_, err := j.redis.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}
