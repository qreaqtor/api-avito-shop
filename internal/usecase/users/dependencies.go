package usersuc

import (
	"context"

	"github.com/qreaqtor/api-avito-shop/internal/models"
)

type repoUser interface {
	GetPassword(ctx context.Context, username string) (string, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username string) (*models.UserRead, error)
	TakeCoin(ctx context.Context, username string, amount uint) error
}

type repoItems interface {
	GetItems(ctx context.Context, username string) ([]*models.InventoryItem, error)
	AddItem(context.Context, *models.Item) error
}

type repoMerch interface {
	GetPrice(ctx context.Context, merch string) (uint, error)
}

type repoTransactions interface {
	GetUserCoinHistory(ctx context.Context, username string) (*models.History, error)
	CreateTransaction(context.Context, *models.Transaction) error
}

type tokenManager interface {
	GenerateToken(username string) (*models.Token, error)
	CheckPassword(hashedPassword, password string) error
	GetHashedPassword(password string) (string, error)
}

type transactionManager interface {
	RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error
}

type UsersDependecnies struct {
	Auth tokenManager

	Tm transactionManager

	Merch        repoMerch
	Users        repoUser
	Items        repoItems
	Transactions repoTransactions
}
