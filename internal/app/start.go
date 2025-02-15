package app

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/qreaqtor/api-avito-shop/internal/api"
	"github.com/qreaqtor/api-avito-shop/internal/config"
	"github.com/qreaqtor/api-avito-shop/internal/lib/auth"
	"github.com/qreaqtor/api-avito-shop/internal/lib/postgres/transactor"
	"github.com/qreaqtor/api-avito-shop/internal/repo/postgres"
	transactionsrepo "github.com/qreaqtor/api-avito-shop/internal/repo/postgres/transactions"
	transrepo "github.com/qreaqtor/api-avito-shop/internal/repo/postgres/transactions"
	usersrepo "github.com/qreaqtor/api-avito-shop/internal/repo/postgres/users"
	usersuc "github.com/qreaqtor/api-avito-shop/internal/usecase/users"
	appserver "github.com/qreaqtor/api-avito-shop/pkg/appServer"
	comlog "github.com/qreaqtor/api-avito-shop/pkg/logging"
)

func StartNewApp(ctx context.Context, cfg config.Config) (*App, error) {
	comlog.SetLogger(cfg.Env)

	router := mux.NewRouter()

	db, err := postgres.GetPostgresConnPool(ctx, cfg.Postgres)
	if err != nil {
		return nil, err
	}

	transactionManager := transactor.NewTransactionManager(db)
	tokenManager := auth.NewTokenManager(cfg.Auth)

	usersrepo := usersrepo.NewUserRepo(transactionManager)
	transactionsRepo := transrepo.NewtransactionsRepo(transactionManager)

	usersUC := usersuc.NewUserUC(usersrepo, tokenManager, transactionsRepo)
	usersApi := api.NewUsersAPI(usersUC)
	usersApi.Register(router)

	appServer := appserver.NewAppServer(ctx, router, cfg.Env, int(cfg.Port)).WithClosers(db)

	app := &App{
		server: appServer,
	}

	err = app.server.Start()
	if err != nil {
		return nil, err
	}

	return app, nil
}
