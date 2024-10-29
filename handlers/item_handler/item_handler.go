package item_handler

import (
	"hkn-be/constants"
	"hkn-be/objects"
	"hkn-be/services"

	"github.com/labstack/echo/v4"
)

type itemHandler struct {
	*services.ServiceCtx
}

// Fungsi DeleteItem untuk menghapus item berdasarkan ID
func (h itemHandler) DeleteItem(c echo.Context) error {
	// Ambil itemId dari URL parameter
	itemIdParam := c.Param("id")

	// Panggil service untuk menghapus item
	err := h.ItemService.DeleteItem(c.Request().Context(), itemIdParam)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	return objects.SetResponse(c, nil, constants.MessageSuccess)
}

// Fungsi CreateItem untuk membuat item baru
func (h itemHandler) CreateItem(c echo.Context) error {
	var newItem *objects.ItemRequest // Misalnya ada struct ItemRequest di objects

	if err := c.Bind(&newItem); err != nil {
		return objects.SetResponse(c, constants.ErrInvalidInput, constants.MessageFailed)
	}

	dtoModels := newItem.ToModel()

	item, err := h.ItemService.CreateItem(c.Request().Context(), dtoModels)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	res := objects.NewItemRes().Map(item)

	return objects.SetResponse(c, nil, res)
}
