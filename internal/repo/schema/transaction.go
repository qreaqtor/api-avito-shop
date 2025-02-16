package schema

import (
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"github.com/uptrace/bun"
)

type TransactionSchema struct {
	bun.BaseModel `bun:"table:transactions"`

	FromUser string `bun:"from_user"`
	ToUser   string `bun:"to_user"`
	Amount   uint   `bun:"amount,notnull"`
}

func NewTransactionSchema(transaction *models.Transaction) *TransactionSchema {
	return &TransactionSchema{
		FromUser: transaction.FromUser,
		ToUser:   transaction.ToUser,
		Amount:   transaction.Amount,
	}
}

func (tr *TransactionSchema) ToDomainTransactionSent() *models.TransactionSent {
	return &models.TransactionSent{
		ToUser: tr.ToUser,
		Amount: tr.Amount,
	}
}

func (tr *TransactionSchema) ToDomainTransactionReceived() *models.TransactionReceived {
	return &models.TransactionReceived{
		FromUser: tr.FromUser,
		Amount:   tr.Amount,
	}
}

type TransactionsSchema []TransactionSchema

func (tr *TransactionsSchema) ToDomainHistory(username string) *models.History {
	history := new(models.History)

	for _, transaction := range *tr {
		if transaction.FromUser == username {
			history.Sent = append(history.Sent, transaction.ToDomainTransactionSent())
			continue
		}
		history.Received = append(history.Received, transaction.ToDomainTransactionReceived())
	}

	return history
}
