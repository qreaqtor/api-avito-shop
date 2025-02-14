package config

import "time"

type AuthConfig struct {
	Secret        string
	TokenLifespan time.Duration ``
}
