package item_handler

import (
	"hkn-be/services"

	"github.com/labstack/echo/v4"
)

type ItemHandlerInterface interface {
	CreateItem(c echo.Context) error
	DeleteItem(c echo.Context) error
	UpdateItem(c echo.Context) error
}

// NewItemHandler menginisialisasi ItemHandler dengan ServiceCtx
func NewItemHandler(ctx *services.ServiceCtx) ItemHandlerInterface {
	return &itemHandler{
		ServiceCtx: ctx,
	}
}
