package user_handler

import (
	"hkn-be/constants"
	"hkn-be/objects"
	"hkn-be/services"

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
	//TODO implement me
	panic("implement me")
}

func (a userHandler) UpdateUser(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
func (a userHandler) GetItemByID(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
func (a userHandler) GetUserByEmail(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}
