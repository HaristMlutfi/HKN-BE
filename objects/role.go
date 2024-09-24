package objects

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id          string
	Code        string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
