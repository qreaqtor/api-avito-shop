package schema

import (
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"github.com/uptrace/bun"
)

type ItemSchema struct {
	bun.BaseModel `bun:"table:inventory_items"`

	Username  string `bun:"username"`
	MerchType string `bun:"merch_type"`
}

func NewItemSchema(item *models.Item) *ItemSchema {
	return &ItemSchema{
		Username:  item.Username,
		MerchType: item.MerchName,
	}
}

type InventoryItemSchema struct {
	bun.BaseModel `bun:"table:inventory_items"`

	MerchType string `bun:"merch_type"`
	Count     int    `bun:"count"`
}

func (i *InventoryItemSchema) ToDomainItem() *models.InventoryItem {
	return &models.InventoryItem{
		ItemType: i.MerchType,
		Quantity: uint(i.Count),
	}
}
