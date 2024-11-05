package models

const TableNameItem = "items"

type Item struct {
	Id         string  `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	NamaBarang string  `gorm:"column:nama_barang;type:varchar(50);not null" json:"nama_barang"`
	Harga      float64 `gorm:"column:harga;type:numeric(12,2);not null" json:"harga"`
	Kategori   string  `gorm:"column:kategori;type:varchar(50)" json:"kategori"`
	Stock      int     `gorm:"column:stock;not null" json:"stock"`
	Discount   int     `gorm:"column:discount;not null" json:"discount"`
}

func (Item) TableName() string {
	return "items"
}
