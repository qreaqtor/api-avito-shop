package models

type TransactionReceived struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type TransactionSent struct {
	ToUser string `json:"toUser" validate:"min=4"`
	Amount int    `json:"amount" validate:"gt=0"`
}
