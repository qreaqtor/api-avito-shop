package itemsrepo

import (
	"context"

	"github.com/qreaqtor/api-avito-shop/internal/lib/postgres/transactor"
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"github.com/qreaqtor/api-avito-shop/internal/repo/schema"
)

type ItemsRepo struct {
	provider transactor.QueryEngineProvider
}

func NewItemsRepo(provider transactor.QueryEngineProvider) *ItemsRepo {
	return &ItemsRepo{
		provider: provider,
	}
}

func (i *ItemsRepo) GetItems(ctx context.Context, username string) ([]*models.InventoryItem, error) {
	db := i.provider.GetQueryEngine(ctx)

	inventory := []schema.InventoryItemSchema{}

	err := db.NewSelect().
		Model((*schema.ItemSchema)(nil)).
		ColumnExpr("merch_type, count(*)").
		Where("username = ?", username).
		Group("merch_type").
		Scan(ctx, &inventory)
	if err != nil {
		return nil, err
	}

	items := make([]*models.InventoryItem, 0, len(inventory))
	for _, item := range inventory {
		items = append(items, item.ToDomainItem())
	}

	return items, nil
}

func (i *ItemsRepo) AddItem(ctx context.Context, item *models.Item) error {
	db := i.provider.GetQueryEngine(ctx)

	itemSchema := schema.NewItemSchema(item)

	_, err := db.NewInsert().
		Model(itemSchema).
		Exec(ctx)
	return err
}
