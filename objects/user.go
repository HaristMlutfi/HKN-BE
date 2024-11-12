package objects

import (
	"time"

	"hkn-be/models"
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

type UserRequest struct {
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
}

func (u *UserRequest) ToDTO() *UserDTO {
	return &UserDTO{
		Password: u.Password,
		Name:     u.Name,
	}
}

type UserDTO struct {
	Password string
	Name     string
}

func (u *UserDTO) ToModel() *models.User {
	return &models.User{
		Password: u.Password,
		Name:     u.Name,
	}
}

func (u *UserDTO) MapFromModel(model *models.User) *UserDTO {
	return &UserDTO{
		Password: model.Password,
		Name:     model.Name,
	}
}

type UserRes struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	//Role     string `json:"role"`
	// Tambahkan field lain sesuai kebutuhan
}

func NewUserRes() UserRes {
	return UserRes{}
}

func (u UserRes) Map(model *models.User) UserRes {
	return UserRes{
		Name:  model.Name,
		Email: model.Email,
		//Role:  model.Role,
		// Mapping field lain sesuai kebutuhan
	}
}
