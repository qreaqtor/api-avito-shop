package models

type History struct {
	Received []TransactionReceived `json:"received"`
	Sent     []TransactionSent     `json:"sent"`
}
