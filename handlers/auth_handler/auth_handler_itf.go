package auth_handler

import (
	"hkn-be/services"

	"github.com/labstack/echo/v4"
)

type AuthHandlerInterface interface {
	Register(c echo.Context) error
	RequestVerificationCode(c echo.Context) error
	VerifyRegistration(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	RefreshToken(c echo.Context) error
}

func NewAuthHandler(ctx *services.ServiceCtx) AuthHandlerInterface {
	return &authHandler{
		ctx,
	}
}
