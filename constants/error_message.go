package constants

import (
	"errors"
)

var (
	ErrOrderOrShippingIdIsNil = errors.New("order or shipping ID is empty")
	ErrPageLimitWrongFormat   = errors.New("\"Page\" or \"Limit\" isn't in number format")
	ErrSaveTokenToRedis       = errors.New("error save token to redis")
	ErrTokenRequired          = errors.New("token is required")
	ErrKeyIsNotInvalidType    = errors.New("key is of invalid type")
	ErrInvalidIssuer          = errors.New("invalid token issuer")
	ErrGetTokenFromRedis      = errors.New("error get token from redis")
	ErrTokenAlreadyExpired    = errors.New("token already expired")
	ErrInvalidEmailAddress    = errors.New("invalid email address")
	ErrPasswordRequired       = errors.New("password is required")
	ErrInvalidJwtToken        = errors.New("invalid token")
	ErrInvalidId              = errors.New("invalid ID")
)
