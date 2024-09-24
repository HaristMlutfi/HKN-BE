package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"hkn-be/handlers"
	"hkn-be/infras/jwt_infra"
)

// CustomValidator is a custom validator for Echo
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates the request payload
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func InitRouter(handler *handlers.HandlerCtx, jwtInterface jwt_infra.JwtInterface) *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowHeaders: []string{
					echo.HeaderOrigin,
					echo.HeaderContentType,
					echo.HeaderAccept,
					echo.HeaderAuthorization,
				},
				AllowMethods: []string{echo.GET, echo.POST, echo.DELETE},
			},
		),
	)
	e.Validator = &CustomValidator{validator: validator.New()}
	setAppRoutes(handler, e, jwtInterface)

	return e
}

func setAppRoutes(handler *handlers.HandlerCtx, e *echo.Echo, jwtInterface jwt_infra.JwtInterface) {
	e.POST("/register", handler.AuthHandler.Register)
	e.POST("/verify-registration", handler.AuthHandler.VerifyRegistration)
	e.POST("/request-verification-code", handler.AuthHandler.RequestVerificationCode)
	e.POST("/login", handler.AuthHandler.Login)

	v1 := e.Group("/v1")
	v1.Use(jwtInterface.GetEchoJwtMiddlewareConfig())

}
