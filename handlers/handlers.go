package handlers

import (
	"hkn-be/handlers/auth_handler"
	"hkn-be/handlers/item_handler"
	"hkn-be/handlers/user_handler"
	"hkn-be/services"
)

type HandlerCtx struct {
	AuthHandler auth_handler.AuthHandlerInterface
	UserHandler user_handler.UserHandlerInterface
	ItemHandler item_handler.ItemHandlerInterface
}

func InitHandlers(ctxServices *services.ServiceCtx) *HandlerCtx {
	authHandler := auth_handler.NewAuthHandler(ctxServices)
	userHandler := user_handler.NewUserHandler(ctxServices)
	itemHandler := item_handler.NewItemHandler(ctxServices)
	return &HandlerCtx{
		AuthHandler: authHandler,
		UserHandler: userHandler,
		ItemHandler: itemHandler,
	}
}
