package objects

import (
	"time"
)

type User struct {
	Id          string
	Email       string
	Password    string
	GoogleId    string
	GoogleToken string
	Name        string
	IsVerified  bool
	Roles       []Role
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
