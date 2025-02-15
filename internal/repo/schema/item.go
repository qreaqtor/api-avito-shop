package schema

import (
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"github.com/uptrace/bun"
)

type ItemSchema struct {
	bun.BaseModel `bun:"table:inventory_items"`

	MerchName string `bun:"merch_name,notnull"`
	Quantity  uint   `bun:"quantity,notnull"`
}

func (i *ItemSchema) ToDomainItem() *models.Item {
	return &models.Item{
		Type: i.MerchName,
		Quantity: i.Quantity,
	}
}
