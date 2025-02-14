package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/qreaqtor/api-avito-shop/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var errConnect = errors.New("can't connect to PostgreSQL")

func GetPostgresConnPool(ctx context.Context, cfg config.PostgresConfig) (*bun.DB, error) {
	connConfig, err := pgx.ParseConfig(cfg.URL)
	if err != nil {
		return nil, err
	}

	sqldb := stdlib.OpenDB(*connConfig)
	err = ping(ctx, cfg, sqldb)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New(),
		maxOpenConnsDB(cfg.MaxOpenConnections),
		maxIdleConnsDB(cfg.MaxIdleConnections),
		maxConnLifetimeDB(cfg.MaxConnLifetime),
		maxConnIdleTimeDB(cfg.MaxConnIdleTime),
	)
	return db, nil
}

func ping(ctx context.Context, cfg config.PostgresConfig, sqldb *sql.DB) error {
	tiker := time.NewTicker(cfg.ConnectionRetryInterval)
	defer tiker.Stop()

RetryLoop:
	for range cfg.ConnectionAttempts {
		select {
		case <-ctx.Done():
			break RetryLoop
		case <-tiker.C:
			err := sqldb.Ping()
			if err == nil {
				return nil
			}
		}
	}

	return errConnect
}
