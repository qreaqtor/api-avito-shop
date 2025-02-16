package schema

import (
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"github.com/uptrace/bun"
)

type UserSchema struct {
	bun.BaseModel `bun:"table:users"`

	Name     string `bun:"username,pk"`
	Password string `bun:"password,notnull"`
	Coins    int    `bun:"coins,default:1000"`
}

func NewUserSchema(user *models.User) *UserSchema {
	return &UserSchema{
		Name:     user.Name,
		Password: user.Password,
		Coins:    int(user.Coins),
	}
}

func (user *UserSchema) ToDomainUser() *models.User {
	return &models.User{
		Name:     user.Name,
		Password: user.Password,
		Coins:    uint(user.Coins),
	}
}
