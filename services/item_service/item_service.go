package item_service

import (
	"context"
	"hkn-be/models"
	"hkn-be/objects"
	"hkn-be/repositories/item_repo"
)

type itemService struct {
	itemRepo item_repo.ItemRepoInterface
}

// CreateItem menambahkan item baru melalui repo
func (s *itemService) CreateItem(ctx context.Context, dto *objects.ItemDto) (*objects.ItemDto, error) {
	item := dto.ToModel()

	res, err := s.itemRepo.CreateItem(ctx, item)

	if err != nil {
		return nil, err
	}

	resDto := dto.MapFromModel(res)

	return resDto, err
}

// GetItemByID mengambil item berdasarkan ID melalui repo
func (s *itemService) GetItemByID(ctx context.Context, id string) (*models.Item, error) {
	return s.itemRepo.GetItemByID(ctx, id)
}

// UpdateItem memperbarui item yang ada melalui repo
func (s *itemService) UpdateItem(ctx context.Context, item *models.Item) error {
	return s.itemRepo.UpdateItem(ctx, item)
}

// DeleteItem menghapus item berdasarkan ID melalui repo
func (s *itemService) DeleteItem(ctx context.Context, id string) error {
	return s.itemRepo.DeleteItem(ctx, id)
}

// ListItems mengambil semua item dari repo
func (s *itemService) ListItems(ctx context.Context) ([]models.Item, error) {
	return s.itemRepo.ListItems(ctx)
}
