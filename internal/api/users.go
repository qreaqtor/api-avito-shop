package api

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/qreaqtor/api-avito-shop/internal/models"
)

const (
	unathorized = "unathorized"
	ok          = "ok"
)

type userService interface {
	CheckAuth(context.Context, *models.AuthInfo) (*models.Token, error)
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

func (u *UsersAPI) Register(r *mux.Router) {
	r.Path("/auth").HandlerFunc(u.auth).Methods(http.MethodPost)
}
