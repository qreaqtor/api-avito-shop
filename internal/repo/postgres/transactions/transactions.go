package transrepo

import (
	"context"

	"github.com/qreaqtor/api-avito-shop/internal/lib/postgres/transactor"
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"github.com/qreaqtor/api-avito-shop/internal/repo/schema"
)

type TransactionsRepo struct {
	provider transactor.QueryEngineProvider
}

func NewTransactionsRepo(provider transactor.QueryEngineProvider) *TransactionsRepo {
	return &TransactionsRepo{
		provider: provider,
	}
}

func (t *TransactionsRepo) GetUserCoinHistory(ctx context.Context, username string) (*models.History, error) {
	db := t.provider.GetQueryEngine(ctx)

	transactionsSchema := schema.TransactionsSchema{}

	sentQuery := db.NewSelect().
		Model((*schema.TransactionSchema)(nil)).
		Where("from_user = ?", username)

	receivedQuery := db.NewSelect().
		Model((*schema.TransactionSchema)(nil)).
		Where("to_user = ?", username)

	err := db.NewSelect().
		With("sent", sentQuery).
		With("received", receivedQuery).
		TableExpr("(SELECT * FROM sent UNION ALL SELECT * FROM received) AS transactions").
		Scan(ctx, &transactionsSchema)
	if err != nil {
		return nil, err
	}

	return transactionsSchema.ToDomainHistory(username), nil
}
