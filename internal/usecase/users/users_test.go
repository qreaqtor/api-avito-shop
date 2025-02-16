package usersuc

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/qreaqtor/api-avito-shop/internal/models"
	repoerr "github.com/qreaqtor/api-avito-shop/internal/repo/err"
	"github.com/stretchr/testify/assert"
)

func TestCheckAuth_UserExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := NewMockrepoUser(ctrl)
	mockAuth := NewMocktokenManager(ctrl)

	uuc := NewUserUC(UsersDependecnies{
		Users: mockUsers,
		Auth:  mockAuth,
	})

	ctx := context.Background()
	authInfo := &models.AuthInfo{
		Username: "testuser",
		Password: "password",
	}
	hashed := "hashedpassword"

	mockUsers.EXPECT().
		GetPassword(ctx, authInfo.Username).
		Return(hashed, nil)

	mockAuth.EXPECT().
		CheckPassword(hashed, authInfo.Password).
		Return(nil)

	tokenMock := &models.Token{
		AuthToken: "validtoken",
	}
	mockAuth.EXPECT().
		GenerateToken(authInfo.Username).
		Return(tokenMock, nil)

	token, err := uuc.CheckAuth(ctx, authInfo)

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, token.AuthToken, token.AuthToken)
}

func TestCheckAuth_UserNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := NewMockrepoUser(ctrl)
	mockAuth := NewMocktokenManager(ctrl)

	uuc := NewUserUC(UsersDependecnies{
		Users: mockUsers,
		Auth:  mockAuth,
	})

	ctx := context.Background()
	authInfo := &models.AuthInfo{
		Username: "newuser",
		Password: "password",
	}
	hashed := "newhashedpassword"

	mockUsers.EXPECT().
		GetPassword(ctx, authInfo.Username).
		Return("", repoerr.ErrNotFound)

	mockAuth.EXPECT().
		GetHashedPassword(authInfo.Password).
		Return(hashed, nil)

	mockUsers.EXPECT().
		CreateUser(ctx, gomock.Any()).
		Return(nil)

	tokenMock := &models.Token{
		AuthToken: "newusertoken",
	}
	mockAuth.EXPECT().
		GenerateToken(authInfo.Username).
		Return(tokenMock, nil)

	token, err := uuc.CheckAuth(ctx, authInfo)

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.Equal(t, tokenMock.AuthToken, token.AuthToken)
}

func TestGetUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := NewMockrepoUser(ctrl)
	mockItems := NewMockrepoItems(ctrl)
	mockTransactions := NewMockrepoTransactions(ctrl)

	uuc := NewUserUC(UsersDependecnies{
		Users:        mockUsers,
		Items:        mockItems,
		Transactions: mockTransactions,
	})

	ctx := context.Background()
	username := "testuser"
	coins := 100
	userRead := &models.UserRead{
		Coins: uint(coins),
	}

	mockUsers.EXPECT().
		GetUser(ctx, username).
		Return(userRead, nil)

	mockItems.EXPECT().
		GetItems(ctx, username).
		Return([]*models.InventoryItem{}, nil)

	mockTransactions.EXPECT().
		GetUserCoinHistory(ctx, username).
		Return(&models.History{}, nil)

	userInfo, err := uuc.GetUserInfo(ctx, username)

	assert.NoError(t, err)
	assert.NotNil(t, userInfo)
	assert.Equal(t, uint(coins), userInfo.Coins)
}

func TestSendCoin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsers := NewMockrepoUser(ctrl)
	mockTransactions := NewMockrepoTransactions(ctrl)
	mockTm := NewMocktransactionManager(ctrl)

	uuc := NewUserUC(UsersDependecnies{
		Users:        mockUsers,
		Transactions: mockTransactions,
		Tm:           mockTm,
	})

	ctx := context.Background()
	transaction := &models.Transaction{
		FromUser: "fromUser",
		ToUser:   "toUser",
		Amount:   100,
	}

	mockTm.EXPECT().RunRepeatableRead(ctx, gomock.Any()).DoAndReturn(
		func(ctxTX context.Context, fn func(ctxTX context.Context) error) error {
			return fn(ctxTX)
		},
	).Return(nil)
	mockUsers.EXPECT().TakeCoin(ctx, transaction.FromUser, transaction.Amount).Return(nil)
	mockTransactions.EXPECT().CreateTransaction(ctx, transaction).Return(nil)

	err := uuc.SendCoin(ctx, transaction)
	assert.NoError(t, err)
}
