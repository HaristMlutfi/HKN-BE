package objects

import "hkn-be/models"

type ItemRequest struct {
	NamaBarang string  `json:"nama_barang" validate:"required"`
	Harga      float64 `json:"harga" validate:"required"`
	Kategori   string  `json:"kategori" validate:"omitempty"`
	Stock      int     `json:"stock" validate:"required"`
	Discount   int     `json:"discount" validate:"omitempty"`
}

func NewItemRequest() *ItemRequest {
	return &ItemRequest{}
}

func (i *ItemRequest) ToModel() *ItemDto {
	return &ItemDto{
		NamaBarang: i.NamaBarang,
		Harga:      i.Harga,
		Kategori:   i.Kategori,
		Stock:      i.Stock,
		Discount:   i.Discount,
	}
}
func (i *ItemRequest) TooModel() *models.Item {
	return &models.Item{
		NamaBarang: i.NamaBarang,
		Harga:      i.Harga,
		Kategori:   i.Kategori,
		Stock:      i.Stock,
		Discount:   i.Discount,
	}
}

type ItemDto struct {
	NamaBarang string
	Harga      float64
	Kategori   string
	Stock      int
	Discount   int
}

func NewItemDto() *ItemDto {
	return &ItemDto{}
}

func (i *ItemDto) ToModel() *models.Item {
	return &models.Item{
		NamaBarang: i.NamaBarang,
		Harga:      i.Harga,
		Kategori:   i.Kategori,
		Stock:      i.Stock,
		Discount:   i.Discount,
	}
}

func (i *ItemDto) MapFromModel(model *models.Item) *ItemDto {
	return &ItemDto{
		NamaBarang: i.NamaBarang,
		Harga:      i.Harga,
		Kategori:   i.Kategori,
		Stock:      i.Stock,
		Discount:   i.Discount,
	}
}

type ItemRes struct {
	NamaBarang string  `json:"nama_barang" validate:"required"`
	Harga      float64 `json:"harga" validate:"required"`
	Kategori   string  `json:"kategori" validate:"omitempty"`
	Stock      int     `json:"stock" validate:"required"`
	Discount   int     `json:"discount" validate:"omitempty"`
}

func NewItemRes() ItemRes {
	return ItemRes{}
}

func (i ItemRes) Map(dto *ItemDto) ItemRes {
	i.NamaBarang = dto.NamaBarang
	i.Harga = dto.Harga
	i.Kategori = dto.Kategori
	i.Stock = dto.Stock
	return i
}
