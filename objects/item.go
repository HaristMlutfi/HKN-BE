package objects

import "hkn-be/models"

type ItemRequest struct {
	NamaBarang string  `json:"nama_barang" validate:"required"`
	Harga      float64 `json:"harga" validate:"required"`
	Kategori   string  `json:"kategori" validate:"omitempty"`
	Stock      int     `json:"stock" validate:"required"`
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
	}
}

type ItemDto struct {
	NamaBarang string
	Harga      float64
	Kategori   string
	Stock      int
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
	}
}

func (i *ItemDto) MapFromModel(model *models.Item) *ItemDto {
	return &ItemDto{
		NamaBarang: i.NamaBarang,
		Harga:      i.Harga,
		Kategori:   i.Kategori,
		Stock:      i.Stock,
	}
}

type ItemRes struct {
	NamaBarang string  `json:"nama_barang" validate:"required"`
	Harga      float64 `json:"harga" validate:"required"`
	Kategori   string  `json:"kategori" validate:"omitempty"`
	Stock      int     `json:"stock" validate:"required"`
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
