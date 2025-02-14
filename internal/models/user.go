package models

type UserInfo struct {
	Coins     uint   `json:"coins"`
	Inventory []Item `json:"inventory"`

	CoinHistory History `json:"coinHistory"`
}

type User struct {
	Name     string
	Password string
	Coins    uint
}
