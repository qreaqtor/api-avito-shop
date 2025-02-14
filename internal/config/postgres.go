package config

import (
	"time"
)

type PostgresConfig struct {
	ConnectionAttempts      int           `yaml:"conn_attempts" env-default:"10"`
	ConnectionRetryInterval time.Duration `yaml:"conn_retry_interval" env-default:"1s"`

	MaxOpenConnections int           `yaml:"max_open_conn" env-default:"50"`
	MaxIdleConnections int           `yaml:"max_idle_conn" env-default:"20"`
	MaxConnLifetime    time.Duration `yaml:"max_conn_lifetime" env-default:"10m"`
	MaxConnIdleTime    time.Duration `yaml:"max_conn_idle_time" env-default:"5m"`

	URL string `env:"DATABASE_URL" env-default:"postgres://user:password@localhost:5432/shop?sslmode=disable"`
}
