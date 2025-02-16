//go:build integration

package integrationstest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sendCoinsRequest struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

func sendCoins(amount int, userTo string, auth *authResponse) (*http.Response, error) {
	url := baseURL + "/sendCoin"

	request := &sendCoinsRequest{
		ToUser: userTo,
		Amount: amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+auth.Token)

	return http.DefaultClient.Do(req)
}

func TestSendCoinsSimple(t *testing.T) {
	userFrom := "userFromSimple"
	userTo := "userToSimple"

	authResp, err := authenticate(userFrom, testPassword)
	assert.NoError(t, err)

	_, err = authenticate(userTo, testPassword)
	assert.NoError(t, err)

	amount := 100

	resp, err := sendCoins(amount, userTo, authResp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
