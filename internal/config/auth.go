package config

import "time"

type AuthConfig struct {
	JWT
	TokenLifespan      time.Duration `yaml:"token_lifespan" env-required:"true" `
	PasswordCostBcrypt int           `yaml:"password_cost_bcrypt" env-required:"true" `
}

type JWT struct {
	SigningKey string `env:"JWT_SIGNING_KEY" env-default:"secret"`
}
