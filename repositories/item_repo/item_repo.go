package item_repo

import (
	"context"
	"hkn-be/models"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type itemRepo struct {
	*gorm.DB
}

// CreateItem menambahkan item baru ke database
func (r *itemRepo) CreateItem(ctx context.Context, item *models.Item) (*models.Item, error) {
	if err := r.WithContext(ctx).Create(item).Error; err != nil {
		log.Error("Failed to create item:", err)
		return nil, err
	}
	return item, nil
}

// GetItemByID mengambil item berdasarkan ID-nya
func (r *itemRepo) GetItemByID(ctx context.Context, id string) (*models.Item, error) {
	var item models.Item
	if err := r.WithContext(ctx).First(&item, "id_barang = ?", id).Error; err != nil {
		log.Error("Failed to get item by ID:", err)
		return nil, err
	}
	return &item, nil
}

// UpdateItem memperbarui item yang sudah ada
func (r *itemRepo) UpdateItem(ctx context.Context, item *models.Item) error {
	if err := r.WithContext(ctx).Save(item).Error; err != nil {
		log.Error("Failed to update item:", err)
		return err
	}
	return nil
}

// DeleteItem menghapus item berdasarkan ID-nya
func (r *itemRepo) DeleteItem(ctx context.Context, id string) error {
	if err := r.WithContext(ctx).Delete(&models.Item{}, "id_barang = ?", id).Error; err != nil {
		log.Error("Failed to delete item:", err)
		return err
	}
	return nil
}

// ListItems mengambil semua item dari database
func (r *itemRepo) ListItems(ctx context.Context) ([]models.Item, error) {
	var items []models.Item
	if err := r.WithContext(ctx).Find(&items).Error; err != nil {
		log.Error("Failed to list items:", err)
		return nil, err
	}
	return items, nil
}
