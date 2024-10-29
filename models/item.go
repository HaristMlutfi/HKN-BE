package models

import "time"

type Item struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Category  string    `gorm:"type:varchar(100);not null"`
	Price     float64   `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (Item) TableName() string {
	return "items"
}
