package models

const TableNameUserRole = "user_roles"

// UserRole mapped from table <user_roles>
type UserRole struct {
	UserId string `gorm:"column:user_id;not null" json:"user_id"`
	RoleId string `gorm:"column:role_id;not null" json:"role_id"`
}
