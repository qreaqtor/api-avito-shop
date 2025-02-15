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

func (t *ItemsRepo) GetItems(ctx context.Context, username string) ([]*models.Item, error) {
	db := t.provider.GetQueryEngine(ctx)

	itemsSchema := []schema.ItemSchema{}

	err := db.NewSelect().
		Model((*schema.ItemSchema)(nil)).
		Where("username = ?", username).
		Scan(ctx, &itemsSchema)
	if err != nil {
		return nil, err
	}

	items := make([]*models.Item, 0, len(itemsSchema))
	for _, item := range itemsSchema {
		items = append(items, item.ToDomainItem())
	}

	return items, nil
}
