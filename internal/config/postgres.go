package config

import (
	"time"
)

type PostgresConfig struct {
	ConnectionAttempts int           `env:"CONNECTION_ATTEMPTS" env-default:"10"`
	ConnectionInterval time.Duration `env:"CONNECTION_INTERVAL" env-default:"1s"`

	URL string `env:"DATABASE_URL" env-default:"postgres://user:password@localhost:5432/shop?sslmode=disable"`
}
