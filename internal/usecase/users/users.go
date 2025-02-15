package usersuc

import (
	"context"
	"errors"

	"github.com/qreaqtor/api-avito-shop/internal/models"
	repoerr "github.com/qreaqtor/api-avito-shop/internal/repo/err"
)

type repoUser interface {
	GetPassword(ctx context.Context, username string) (string, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, username string) (*models.UserRead, error)
}

type repoItems interface {
	GetItems(ctx context.Context, username string) ([]*models.Item, error)
}

type repoTransactions interface {
	GetUserCoinHistory(ctx context.Context, username string) (*models.History, error)
}

type tokenManager interface {
	GenerateToken(username string) (*models.Token, error)
	CheckPassword(hashedPassword, password string) error
	GetHashedPassword(password string) (string, error)
}

type UserUC struct {
	users        repoUser
	auth         tokenManager
	items        repoItems
	transactions repoTransactions
}

func NewUserUC(users repoUser, auth tokenManager, items repoItems, transactions repoTransactions) *UserUC {
	return &UserUC{
		users: users,
		auth:  auth,
	}
}

func (u *UserUC) CheckAuth(ctx context.Context, auth *models.AuthInfo) (*models.Token, error) {
	hashedPassword, err := u.users.GetPassword(ctx, auth.Username)
	userExists := !errors.Is(err, repoerr.ErrNotFound)

	if err != nil && userExists {
		return nil, err
	}

	if userExists {
		err = u.auth.CheckPassword(hashedPassword, auth.Password)
	} else {
		err = u.createUser(ctx, auth)
	}
	if err != nil {
		return nil, err
	}

	return u.auth.GenerateToken(auth.Username)
}

func (u *UserUC) createUser(ctx context.Context, auth *models.AuthInfo) error {
	hashedPassword, err := u.auth.GetHashedPassword(auth.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     auth.Username,
		Password: hashedPassword,
	}

	return u.users.CreateUser(ctx, user)
}

func (u *UserUC) GetUser(ctx context.Context, username string) (*models.UserInfo, error) {
	userRead, err := u.users.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}

	items, err := u.items.GetItems(ctx, username)
	if err != nil {
		return nil, err
	}

	history, err := u.transactions.GetUserCoinHistory(ctx, username)
	if err != nil {
		return nil, err
	}

	userInfo := &models.UserInfo{
		Coins:       userRead.Coins,
		Inventory:   items,
		CoinHistory: history,
	}
	return userInfo, nil
}
