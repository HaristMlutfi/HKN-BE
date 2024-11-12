package item_handler

import (
	"hkn-be/constants"
	"hkn-be/objects"
	"hkn-be/services"

	"fmt"

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
	fmt.Println("Request received") // Tambahkan log ini
	var newItem *objects.ItemRequest

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

func (h itemHandler) UpdateItem(c echo.Context) error {
	// Ambil itemId dari URL parameter
	itemIdParam := c.Param("id")

	// Ambil request body dan bind ke ItemRequest
	var req objects.ItemRequest
	if err := c.Bind(&req); err != nil {
		return objects.SetResponse(c, constants.ErrInvalidInput, constants.MessageFailed)
	}

	// Konversi ItemRequest ke models.Item
	itemToUpdate := req.TooModel()
	itemToUpdate.Id = itemIdParam // Assign ID yang diterima dari URL parameter

	// Panggil service untuk update item
	err := h.ItemService.UpdateItem(c.Request().Context(), itemToUpdate)
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	return objects.SetResponse(c, nil, constants.MessageSuccess)
}

func (h *itemHandler) ListItems(c echo.Context) error {
	// Panggil fungsi di service untuk mendapatkan daftar item
	items, err := h.ItemService.ListItems(c.Request().Context())
	if err != nil {
		return objects.SetResponse(c, err, constants.MessageFailed)
	}

	// Konversi items ke response format yang sesuai
	var itemResponses []objects.ItemRes
	for _, item := range items {
		itemResponse := objects.NewItemRes().Map(&objects.ItemDto{
			NamaBarang: item.NamaBarang,
			Harga:      item.Harga,
			Kategori:   item.Kategori,
			Stock:      item.Stock,
			Discount:   item.Discount,
		})
		itemResponses = append(itemResponses, itemResponse)
	}

	// Kirimkan itemResponses sebagai response sukses
	return objects.SetResponse(c, nil, itemResponses)
}
