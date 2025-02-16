package usersuc

import (
	"context"
	"errors"

	"github.com/qreaqtor/api-avito-shop/internal/models"
	repoerr "github.com/qreaqtor/api-avito-shop/internal/repo/err"
)

var (
	errBadTransaction = errors.New("bad transaction")
)

type UserUC struct {
	auth tokenManager

	tm transactionManager

	merch        repoMerch
	users        repoUser
	items        repoItems
	transactions repoTransactions
}

func NewUserUC(deps UsersDependecnies) *UserUC {
	return &UserUC{
		users:        deps.Users,
		auth:         deps.Auth,
		items:        deps.Items,
		transactions: deps.Transactions,
		tm:           deps.Tm,
		merch:        deps.Merch,
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

func (u *UserUC) GetUserInfo(ctx context.Context, username string) (*models.UserInfo, error) {
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

func (u *UserUC) SendCoin(ctx context.Context, transaction *models.Transaction) error {
	if transaction.FromUser == transaction.ToUser {
		return errBadTransaction
	}
	return u.tm.RunRepeatableRead(ctx, func(ctxTX context.Context) error {
		err := u.users.TakeCoin(ctx, transaction.FromUser, transaction.Amount)
		if err != nil {
			return err
		}

		return u.transactions.CreateTransaction(ctx, transaction)
	})
}

func (u *UserUC) BuyItem(ctx context.Context, item *models.Item) error {
	return u.tm.RunRepeatableRead(ctx, func(ctxTX context.Context) error {
		merchPrice, err := u.merch.GetPrice(ctx, item.MerchName)
		if err != nil {
			return err
		}

		err = u.users.TakeCoin(ctx, item.Username, merchPrice)
		if err != nil {
			return err
		}

		return u.items.AddItem(ctx, item)
	})
}
