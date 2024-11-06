package user_handler

import (
	"hkn-be/constants"
	"hkn-be/objects"
	"hkn-be/services"

	"hkn-be/models"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	*services.ServiceCtx
}

func (a userHandler) DeleteUser(c echo.Context) error {
	// Ambil userId dari URL parameter
	userIdParam := c.Param("id")

	// Panggil service untuk menghapus user
	err := a.UserService.DeleteUser(c.Request().Context(), userIdParam)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	return objects.SetResponse(c, nil, constants.MessageSuccess)
}

func (a userHandler) CreateUser(c echo.Context) error {
	// Buat objek user request dari body
	var req objects.User
	if err := c.Bind(&req); err != nil {
		return objects.SetResponse(c, constants.ErrInvalidInput, constants.MessageFailed)
	}

	// Validasi input
	if err := c.Validate(req); err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	// Panggil service untuk membuat user baru
	user := models.User{
		Email:      req.Email,
		Password:   req.Password,
		Name:       req.Name,
		IsVerified: req.IsVerified,
	}

	err := a.UserService.CreateUser(c.Request().Context(), user)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	return objects.SetResponse(c, nil, constants.MessageSuccess)
}

// file: handlers/user_handler.go
// file: handlers/user_handler.go
func (h userHandler) UpdateUser(c echo.Context) error {
	// Ambil userId dari URL parameter
	userIdParam := c.Param("id")

	// Ambil request body
	var req objects.UserRequest
	if err := c.Bind(&req); err != nil {
		return objects.SetResponse(c, constants.ErrInvalidInput, constants.MessageFailed)
	}

	// Konversi dari UserRequest ke UserDTO
	userDTO := req.ToDTO()

	// Panggil service untuk mengupdate user
	err := h.UserService.UpdateUser(c.Request().Context(), userIdParam, userDTO.ToModel())
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	return objects.SetResponse(c, nil, constants.MessageSuccess)
}

func (a userHandler) GetItemByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
func (a userHandler) GetUserByEmail(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
