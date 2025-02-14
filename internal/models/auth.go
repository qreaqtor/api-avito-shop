package models

type AuthInfo struct {
	Username string `json:"username" validate:"min=4"`
	Password string `json:"password" validate:"min=4"`
}
