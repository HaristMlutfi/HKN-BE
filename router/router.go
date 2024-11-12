package router

import (
	"hkn-be/handlers"
	"hkn-be/infras/jwt_infra"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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
	//Auth routes
	authRoutes := e.Group("/auth")
	authRoutes.POST("/register", handler.AuthHandler.Register)
	authRoutes.POST("/verify-registration", handler.AuthHandler.VerifyRegistration)
	authRoutes.POST("/request-verification-code", handler.AuthHandler.RequestVerificationCode)
	authRoutes.POST("/login", handler.AuthHandler.Login)
	//User routes
	userRoutes := e.Group("/user")
	userRoutes.DELETE("/delete/:id", handler.UserHandler.DeleteUser)
	userRoutes.POST("/create-user", handler.UserHandler.CreateUser)
	userRoutes.PUT("/update-user/:id", handler.UserHandler.UpdateUser)
	userRoutes.GET("/get-user/:id", handler.UserHandler.GetUserById)
	userRoutes.GET("/get-user-by-email/:email", handler.UserHandler.GetUserByEmail)

	//Item routes
	itemRoutes := e.Group("/item")
	itemRoutes.GET("/list-items", handler.ItemHandler.ListItems)
	itemRoutes.PUT("/update-item/:id", handler.ItemHandler.UpdateItem)
	itemRoutes.POST("/create-item", handler.ItemHandler.CreateItem)
	itemRoutes.DELETE("/delete-item/:id", handler.ItemHandler.DeleteItem)

	v1 := e.Group("/v1")
	v1.Use(jwtInterface.GetEchoJwtMiddlewareConfig())

}
