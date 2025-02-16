package merchrepo

import (
	"context"

	"github.com/qreaqtor/api-avito-shop/internal/lib/postgres/transactor"
	"github.com/qreaqtor/api-avito-shop/internal/repo/schema"
)

type MerchRepo struct {
	provider transactor.QueryEngineProvider
}

func NewMerchRepo(provider transactor.QueryEngineProvider) *MerchRepo {
	return &MerchRepo{
		provider: provider,
	}
}

func (m *MerchRepo) GetPrice(ctx context.Context, merch string) (uint, error) {
	db := m.provider.GetQueryEngine(ctx)

	var price uint
	err := db.NewSelect().
		Model((*schema.MerchSchema)(nil)).
		Column("price").
		Where("merch_type = ?", merch).
		Scan(ctx, &price)
	if err != nil {
		return 0, err
	}
	return price, nil
}
