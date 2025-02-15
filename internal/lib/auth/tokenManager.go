package auth

import "github.com/qreaqtor/api-avito-shop/internal/config"

type TokenManager struct {
	cfg config.AuthConfig
}

func NewTokenManager(cfg config.AuthConfig) *TokenManager {
	return &TokenManager{
		cfg: cfg,
	}
}
