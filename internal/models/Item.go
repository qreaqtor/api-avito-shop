package models

type InventoryItem struct {
	ItemType string `json:"type"`
	Quantity uint   `json:"quantity"`
}

type Item struct {
	Username  string
	MerchName string
}
