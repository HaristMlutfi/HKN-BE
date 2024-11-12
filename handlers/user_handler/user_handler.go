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

func (h *userHandler) GetUserById(c echo.Context) error {
	// Ambil userId dari parameter URL
	userId := c.Param("id")

	// Panggil service untuk mendapatkan data user berdasarkan userId
	user, err := h.UserService.GetUserById(c.Request().Context(), userId)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	// Konversi data user ke response jika menggunakan DTO atau struct respons tertentu
	userResponse := objects.NewUserRes().Map(user)

	// Kembalikan respons dengan data userResponse
	return objects.SetResponse(c, nil, userResponse)
}

func (a userHandler) GetUserByEmail(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
