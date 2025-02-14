package postgres

import (
	"time"

	"github.com/uptrace/bun"
)

func maxOpenConnsDB(maxConns int) bun.DBOption {
	return func(db *bun.DB) {
		db.SetMaxOpenConns(maxConns)
	}
}

func maxIdleConnsDB(maxConns int) bun.DBOption {
	return func(db *bun.DB) {
		db.SetMaxIdleConns(maxConns)
	}
}

func maxConnLifetimeDB(maxLifetime time.Duration) bun.DBOption {
	return func(db *bun.DB) {
		db.SetConnMaxLifetime(maxLifetime)
	}
}

func maxConnIdleTimeDB(maxIdletime time.Duration) bun.DBOption {
	return func(db *bun.DB) {
		db.SetConnMaxIdleTime(maxIdletime)
	}
}
