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
}

type tokenManager interface {
	GenerateToken(username string) (*models.Token, error)
	CheckPassword(hashedPassword, password string) error
	GetHashedPassword(password string) (string, error)
}

type UserUC struct {
	users repoUser

	auth tokenManager
}

func NewUserUC() *UserUC {
	return &UserUC{}
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
