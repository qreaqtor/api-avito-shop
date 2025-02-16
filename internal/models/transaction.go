package models

type TransactionReceived struct {
	FromUser string `json:"fromUser"`
	Amount   uint   `json:"amount"`
}

type TransactionSent struct {
	ToUser string `json:"toUser" validate:"min=4"`
	Amount uint   `json:"amount" validate:"gt=0"`
}

type Transaction struct {
	FromUser string
	ToUser   string
	Amount   uint
}
