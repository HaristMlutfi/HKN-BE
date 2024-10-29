package item_service

import (
	"context"
	"hkn-be/models"
	"hkn-be/objects"
	"hkn-be/repositories/item_repo"
)

type ItemServiceInterface interface {
	CreateItem(ctx context.Context, item *objects.ItemDto) (*objects.ItemDto, error)
	GetItemByID(ctx context.Context, id string) (*models.Item, error)
	UpdateItem(ctx context.Context, item *models.Item) error
	DeleteItem(ctx context.Context, id string) error
	ListItems(ctx context.Context) ([]models.Item, error)
}

func NewItemService(itemRepo item_repo.ItemRepoInterface) ItemServiceInterface {
	return &itemService{
		itemRepo: itemRepo,
	}
}
