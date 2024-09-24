package role_repo

import (
	"context"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"hkn-be/constants"
	"hkn-be/models"
)

type roleRepo struct {
	*gorm.DB
}

func (r roleRepo) GetRoleById(ctx context.Context, roleId string) (*models.Role, error) {
	var result models.Role
	tx := r.DB.WithContext(ctx).First(&result, constants.QueryById, roleId)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, tx.Error
	}
	return &result, nil
}

func (r roleRepo) GetRoleByCode(ctx context.Context, code string) (*models.Role, error) {
	var result models.Role
	tx := r.DB.WithContext(ctx).First(&result, "code = ?", code)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, tx.Error
	}
	return &result, nil
}
