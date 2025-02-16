package api

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/qreaqtor/api-avito-shop/internal/api/middlewares"
	"github.com/qreaqtor/api-avito-shop/internal/models"
)

type userService interface {
	CheckAuth(context.Context, *models.AuthInfo) (*models.Token, error)
	GetUserInfo(context.Context, string) (*models.UserInfo, error)
	SendCoin(context.Context, *models.Transaction) error
	BuyItem(context.Context, *models.Item) error
}

type UsersAPI struct {
	valid *validator.Validate

	users userService
}

func NewUsersAPI(users userService) *UsersAPI {
	return &UsersAPI{
		users: users,
		valid: validator.New(),
	}
}

func (u *UsersAPI) Register(r *mux.Router, auth middlewares.TokenManager) {
	authMiddleware := middlewares.NewAuthMiddleware(auth)

	r.Path("/auth").
		HandlerFunc(u.auth).
		Methods(http.MethodPost)

	r.Path("/sendCoin").
		Handler(authMiddleware.NextFunc(u.sendCoin)).
		Methods(http.MethodPost)

	r.Path("/info").
		Handler(authMiddleware.NextFunc(u.info)).
		Methods(http.MethodGet)

	r.Path("/buy/{item}").
		Handler(authMiddleware.NextFunc(u.buyItem)).
		Methods(http.MethodGet)
}
