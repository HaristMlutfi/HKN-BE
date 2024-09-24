package models

import (
	"gorm.io/gorm"
	"time"
)

const TableNameUser = "users"

type User struct {
	Id         string         `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Email      string         `gorm:"column:email;type:character varying(50);not null" json:"email"`
	Password   string         `gorm:"column:password;type:character varying;not null" json:"password"`
	GoogleId   *string        `gorm:"column:google_id;type:character varying(50);null" json:"google_id"`
	Name       string         `gorm:"column:name;type:character varying(100);not null" json:"name"`
	IsVerified bool           `gorm:"column:is_verified;type:boolean;not null" json:"is_verified"`
	CreatedAt  *time.Time     `gorm:"column:created_at;type:timestamp without time zone;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  *time.Time     `gorm:"column:updated_at;type:timestamp without time zone" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp without time zone" json:"deleted_at"`
}

func (*User) TableName() string {
	return TableNameUser
}
