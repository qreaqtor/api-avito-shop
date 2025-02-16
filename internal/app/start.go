package app

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/qreaqtor/api-avito-shop/internal/api"
	"github.com/qreaqtor/api-avito-shop/internal/config"
	"github.com/qreaqtor/api-avito-shop/internal/lib/auth"
	"github.com/qreaqtor/api-avito-shop/internal/lib/postgres/transactor"
	"github.com/qreaqtor/api-avito-shop/internal/repo/postgres"
	itemsrepo "github.com/qreaqtor/api-avito-shop/internal/repo/postgres/items"
	merchrepo "github.com/qreaqtor/api-avito-shop/internal/repo/postgres/merch"
	transrepo "github.com/qreaqtor/api-avito-shop/internal/repo/postgres/transactions"
	usersrepo "github.com/qreaqtor/api-avito-shop/internal/repo/postgres/users"
	usersuc "github.com/qreaqtor/api-avito-shop/internal/usecase/users"
	appserver "github.com/qreaqtor/api-avito-shop/pkg/appServer"
	comlog "github.com/qreaqtor/api-avito-shop/pkg/logging"
)

func StartNewApp(ctx context.Context, cfg config.Config) (*App, error) {
	comlog.SetLogger(cfg.Env)

	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	db, err := postgres.GetPostgresConnPool(ctx, cfg.Postgres)
	if err != nil {
		return nil, err
	}

	transactionManager := transactor.NewTransactionManager(db)
	tokenManager := auth.NewTokenManager(cfg.Auth)

	usersRepo := usersrepo.NewUserRepo(transactionManager)
	transactionsRepo := transrepo.NewTransactionsRepo(transactionManager)
	itemsRepo := itemsrepo.NewItemsRepo(transactionManager)
	merchRepo := merchrepo.NewMerchRepo(transactionManager)

	deps := usersuc.UsersDependecnies{
		Auth:         tokenManager,
		Tm:           transactionManager,
		Users:        usersRepo,
		Transactions: transactionsRepo,
		Items:        itemsRepo,
		Merch:        merchRepo,
	}

	usersUC := usersuc.NewUserUC(deps)
	usersApi := api.NewUsersAPI(usersUC)
	usersApi.Register(router, tokenManager)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)(router)

	appServer := appserver.NewAppServer(ctx, corsHandler, cfg.Env, int(cfg.Port)).WithClosers(db)

	app := &App{
		server: appServer,
	}

	err = app.server.Start()
	if err != nil {
		return nil, err
	}

	return app, nil
}
