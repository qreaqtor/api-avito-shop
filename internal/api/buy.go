package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qreaqtor/api-avito-shop/internal/lib/auth"
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"github.com/qreaqtor/api-avito-shop/pkg/httprocess"
	logmsg "github.com/qreaqtor/api-avito-shop/pkg/logging/message"
)

const (
	itemPathParam = "item"

	noItem = "path should contains item type"
)

func (u *UsersAPI) buyItem(w http.ResponseWriter, r *http.Request) {
	logMsg := logmsg.NewLogMsg(r.RequestURI, r.Method)

	itemName, ok := mux.Vars(r)[itemPathParam]
	if !ok || itemName == "" {
		httprocess.WriteError(w, logMsg.WithText(noItem).WithStatus(http.StatusBadRequest))
		return
	}

	username, err := auth.ExtractUsername(r.Context())
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(unathorized).WithStatus(http.StatusUnauthorized))
		return
	}

	item := newItem(username, itemName)
	err = u.users.BuyItem(r.Context(), item)
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(err.Error()).WithStatus(http.StatusBadRequest))
		return
	}

	httprocess.WriteData(w,
		logMsg.WithText(success).WithStatus(http.StatusOK),
		success,
	)
}

func newItem(username, item string) *models.Item {
	return &models.Item{
		Username:  username,
		MerchName: item,
	}
}
