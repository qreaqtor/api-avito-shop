package api

import (
	"net/http"

	"github.com/qreaqtor/api-avito-shop/internal/lib/auth"
	"github.com/qreaqtor/api-avito-shop/pkg/httprocess"
	logmsg "github.com/qreaqtor/api-avito-shop/pkg/logging/message"
)

const success = "success"

func (u *UsersAPI) info(w http.ResponseWriter, r *http.Request) {
	logMsg := logmsg.NewLogMsg(r.RequestURI, r.Method)

	username, err := auth.ExtractUsername(r.Context())
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(err.Error()).WithStatus(http.StatusUnauthorized))
		return
	}

	userInfo, err := u.users.GetUserInfo(r.Context(), username)
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(err.Error()).WithStatus(http.StatusBadRequest))
		return
	}

	httprocess.WriteData(w,
		logMsg.WithText(success).WithStatus(http.StatusOK),
		userInfo,
	)
}
