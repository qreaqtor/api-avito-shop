package transactor

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type transactionKey string

const key = transactionKey("tx")

type (
	QueryEngineProvider interface {
		GetQueryEngine(ctx context.Context) bun.IDB // tx OR pool
	}

	TransactionManager struct {
		db *bun.DB
	}
)

func NewTransactionManager(db *bun.DB) *TransactionManager {
	return &TransactionManager{db}
}

func (tm *TransactionManager) RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error {
	txOptions := &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	}
	return tm.db.RunInTx(ctx, txOptions, func(ctx context.Context, tx bun.Tx) error {
		return fx(context.WithValue(ctx, key, tx))
	})
}

func (tm *TransactionManager) GetQueryEngine(ctx context.Context) bun.IDB {
	tx, ok := ctx.Value(key).(bun.IDB)
	if ok && tx != nil {
		return tx
	}

	return tm.db
}
