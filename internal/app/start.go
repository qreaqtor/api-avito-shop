package app

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/qreaqtor/api-avito-shop/internal/config"
	"github.com/qreaqtor/api-avito-shop/internal/repo/postgres"
	appserver "github.com/qreaqtor/api-avito-shop/pkg/appServer"
	comlog "github.com/qreaqtor/api-avito-shop/pkg/logging"
)

func StartNewApp(ctx context.Context, cfg config.Config) (*App, error) {
	comlog.SetLogger(cfg.Env)

	router := mux.NewRouter()

	connPool, err := postgres.GetPostgresConnPool(ctx, cfg.Postgres)
	if err != nil {
		return nil, err
	}

	//repo := postgres.NewContainerRepo(conn, cfg.UpdatedPeriod)
	//uc := usecase.NewContainerUC(ctx, repo, cfg.WsWritePeriod)
	//api := api.NewContainersAPI(uc)
	//api.Register(router)

	appServer := appserver.NewAppServer(ctx, router, cfg.Env, int(cfg.Port)).WithClosers(connPool)

	app := &App{
		server: appServer,
	}

	err = app.server.Start()
	if err != nil {
		return nil, err
	}

	return app, nil
}
