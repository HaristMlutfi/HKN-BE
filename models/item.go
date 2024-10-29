package models

import (
	"github.com/google/uuid"
)

type Item struct {
	IDBarang   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id_barang"`
	NamaBarang string    `gorm:"type:varchar(50);unique;not null" json:"nama_barang"`
	Harga      float64   `gorm:"type:numeric(12,2);not null" json:"harga"`
	Kategori   string    `gorm:"type:varchar(50)" json:"kategori"`
	Stock      int       `gorm:"not null" json:"stock"`
}

func (Item) TableName() string {
	return "items"
}
