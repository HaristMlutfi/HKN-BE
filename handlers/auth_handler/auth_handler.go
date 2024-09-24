package auth_handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"hkn-be/constants"
	"hkn-be/infras/jwt_infra"
	"hkn-be/objects"
	"hkn-be/services"
)

type authHandler struct {
	*services.ServiceCtx
}

func (a authHandler) Register(c echo.Context) error {
	var request objects.RegistrationRequest
	err := c.Bind(&request)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}
	// Validate request payload
	err = c.Validate(request)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	err = a.AuthService.Register(
		c.Request().Context(), objects.User{
			Email:    request.Email,
			Password: request.Password,
			Name:     request.Name,
		},
	)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	return objects.SetResponse(c, err, constants.MessageSuccess)
}

func (a authHandler) RequestVerificationCode(c echo.Context) error {
	var request objects.RequestVerificationCodeRequest
	err := c.Bind(&request)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	// Validate request payload
	err = c.Validate(request)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	err = a.AuthService.RequestVerificationCode(c.Request().Context(), request.Email)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	return objects.SetResponse(c, err, constants.MessageSuccess)
}

func (a authHandler) VerifyRegistration(c echo.Context) error {
	var request objects.VerifyRegistrationRequest
	err := c.Bind(&request)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}
	// Validate request payload
	err = c.Validate(request)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	err = a.AuthService.VerifyUserRegistration(c.Request().Context(), request)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	return objects.SetResponse(c, err, constants.MessageSuccess)
}

func (a authHandler) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	res, err := a.AuthService.Login(
		c.Request().Context(), objects.LoginRequest{
			Email:    email,
			Password: password,
		},
	)
	return objects.SetResponse(c, err, res)
}

func (a authHandler) Logout(c echo.Context) error {
	jwtClaim := c.Get("authKey").(*jwt.Token).Claims.(*jwt_infra.JwtCustomClaims)
	err := a.AuthService.Logout(c.Request().Context(), jwtClaim.Id)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	} else {
		return objects.SetResponse(c, err, constants.MessageSuccess)
	}

}

func (a authHandler) RefreshToken(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
