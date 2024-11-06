package user_handler

import (
	"hkn-be/services"

	"github.com/labstack/echo/v4"
)

type UserHandlerInterface interface {
	DeleteUser(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
}

func NewUserHandler(ctx *services.ServiceCtx) UserHandlerInterface {
	return &userHandler{
		ServiceCtx: ctx,
	}
}
