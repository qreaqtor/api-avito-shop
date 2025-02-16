package api

import (
	"net/http"

	"github.com/qreaqtor/api-avito-shop/internal/lib/auth"
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"github.com/qreaqtor/api-avito-shop/pkg/httprocess"
	logmsg "github.com/qreaqtor/api-avito-shop/pkg/logging/message"
)

func (u *UsersAPI) sendCoin(w http.ResponseWriter, r *http.Request) {
	logMsg := logmsg.NewLogMsg(r.RequestURI, r.Method)

	transactionSent := new(models.TransactionSent)
	err := httprocess.ReadRequestBody(r, transactionSent)
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(err.Error()).WithStatus(http.StatusBadRequest))
		return
	}

	err = u.valid.StructCtx(r.Context(), transactionSent)
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(err.Error()).WithStatus(http.StatusUnprocessableEntity))
		return
	}

	username, err := auth.ExtractUsername(r.Context())
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(unathorized).WithStatus(http.StatusUnauthorized))
		return
	}

	transaction := newTransaction(transactionSent, username)
	err = u.users.SendCoin(r.Context(), transaction)
	if err != nil {
		httprocess.WriteError(w, logMsg.WithText(err.Error()).WithStatus(http.StatusBadRequest))
		return
	}

	httprocess.WriteData(w,
		logMsg.WithText(success).WithStatus(http.StatusOK),
		success,
	)
}

func newTransaction(transaction *models.TransactionSent, fromUser string) *models.Transaction {
	return &models.Transaction{
		FromUser: fromUser,
		ToUser:   transaction.ToUser,
		Amount:   transaction.Amount,
	}
}
