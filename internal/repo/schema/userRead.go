package schema

import (
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"github.com/uptrace/bun"
)

type UserReadSchema struct {
	bun.BaseModel `bun:"table:users"`

	Coins int `bun:"coins,notnull"`
}

func (user *UserReadSchema) ToDomainUserRead() *models.UserRead {
	return &models.UserRead{
		Coins: uint(user.Coins),
	}
}
