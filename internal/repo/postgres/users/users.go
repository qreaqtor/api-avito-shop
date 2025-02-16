package usersrepo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/qreaqtor/api-avito-shop/internal/lib/postgres/transactor"
	"github.com/qreaqtor/api-avito-shop/internal/models"
	repoerr "github.com/qreaqtor/api-avito-shop/internal/repo/err"
	"github.com/qreaqtor/api-avito-shop/internal/repo/schema"
)

type UsersRepo struct {
	provider transactor.QueryEngineProvider
}

func NewUserRepo(provider transactor.QueryEngineProvider) *UsersRepo {
	return &UsersRepo{
		provider: provider,
	}
}

func (u *UsersRepo) CreateUser(ctx context.Context, user *models.User) error {
	db := u.provider.GetQueryEngine(ctx)

	userSchema := schema.NewUserSchema(user)

	_, err := db.NewInsert().
		Model(userSchema).
		Exec(ctx)
	return err
}

func (u *UsersRepo) GetPassword(ctx context.Context, username string) (string, error) {
	db := u.provider.GetQueryEngine(ctx)

	var password string
	err := db.NewSelect().
		Model((*schema.UserSchema)(nil)).
		Column("password").
		Where("username = ?", username).
		Scan(ctx, &password)
	if errors.Is(err, sql.ErrNoRows) {
		return "", repoerr.ErrNotFound
	}

	return password, err
}

func (u *UsersRepo) GetUser(ctx context.Context, username string) (*models.UserRead, error) {
	db := u.provider.GetQueryEngine(ctx)

	userReadSchema := new(schema.UserReadSchema)

	err := db.NewSelect().
		Model(userReadSchema).
		Where("username = ?", username).
		Scan(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repoerr.ErrNotFound
	}

	return userReadSchema.ToDomainUserRead(), err
}

func (u *UsersRepo) TakeCoin(ctx context.Context, username string, amount uint) error {
	db := u.provider.GetQueryEngine(ctx)

	_, err := db.NewUpdate().
		Model((*schema.UserReadSchema)(nil)).
		Set("coins = coins - ?", amount).
		Where("username = ?", username).
		Exec(ctx)
	return err
}
