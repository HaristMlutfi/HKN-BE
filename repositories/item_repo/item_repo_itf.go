package item_repo

import (
	"context"
	"hkn-be/models"

	"gorm.io/gorm"
)

type ItemRepoInterface interface {
	CreateItem(ctx context.Context, item *models.Item) (*models.Item, error)
	GetItemByID(ctx context.Context, id string) (*models.Item, error)
	UpdateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, id string) error
	ListItems(ctx context.Context) ([]models.Item, error)
}

func NewItemRepo(db *gorm.DB) ItemRepoInterface {
	return &itemRepo{
		db,
	}
}
