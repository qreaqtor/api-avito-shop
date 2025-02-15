package api

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"github.com/qreaqtor/api-avito-shop/pkg/httprocess"
	logmsg "github.com/qreaqtor/api-avito-shop/pkg/logging/message"
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
	}
}

func (u *UsersAPI) Register(r *mux.Router) {
	r.Path("/auth").HandlerFunc(u.auth).Methods(http.MethodPost)
}

func (u *UsersAPI) auth(w http.ResponseWriter, r *http.Request) {
	logMsg := logmsg.NewLogMsg(r.RequestURI, r.Method)

	authRequest := new(models.AuthInfo)
	err := httprocess.ReadRequestBody(r, authRequest)
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(err.Error()).WithStatus(http.StatusBadRequest))
		return
	}

	err = u.valid.StructCtx(r.Context(), authRequest)
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(err.Error()).WithStatus(http.StatusUnprocessableEntity))
		return
	}

	token, err := u.users.CheckAuth(r.Context(), authRequest)
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(unathorized).WithStatus(http.StatusUnauthorized))
		return
	}

	httprocess.WriteData(w,
		logMsg.WithText(ok).WithStatus(http.StatusOK),
		token,
	)
}
