package handlers

import (
	"hkn-be/handlers/auth_handler"
	"hkn-be/services"
)

type HandlerCtx struct {
	AuthHandler auth_handler.AuthHandlerInterface
}

func InitHandlers(ctxServices *services.ServiceCtx) *HandlerCtx {
	authHandler := auth_handler.NewAuthHandler(ctxServices)
	return &HandlerCtx{
		AuthHandler: authHandler,
	}
}
