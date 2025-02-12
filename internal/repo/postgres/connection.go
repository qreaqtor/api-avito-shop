package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/qreaqtor/api-avito-shop/internal/config"
)

var errConnect = errors.New("can't connect to PostgreSQL")

func GetPostgresConnPool(ctx context.Context, cfg config.PostgresConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, cfg.URL)
	if err != nil {
		return nil, err
	}

	tiker := time.NewTicker(cfg.ConnectionInterval)
	defer tiker.Stop()

RetryLoop:
	for range cfg.ConnectionAttempts {
		select {
		case <-ctx.Done():
			break RetryLoop
		case <-tiker.C:
			err = pool.Ping(ctx)
			if err == nil {
				return pool, nil
			}
		}
	}

	return nil, errConnect
}
