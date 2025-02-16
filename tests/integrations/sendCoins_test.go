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

// func TestSendCoinsWhenUserToNotExists(t *testing.T) {
// 	userFrom := "userFromExists"
// 	userTo := "userToNotExists"

// 	authResp, err := authenticate(userFrom, testPassword)
// 	assert.NoError(t, err)

// 	amount := 100

// 	resp, err := sendCoins(amount, userTo, authResp)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
// }

// func TestSendCoinsToSelf(t *testing.T) {
// 	userFrom := "userFromSelf"
// 	userTo := userFrom

// 	authResp, err := authenticate(userFrom, testPassword)
// 	assert.NoError(t, err)

// 	_, err = authenticate(userTo, testPassword)
// 	assert.NoError(t, err)

// 	amount := 100

// 	resp, err := sendCoins(amount, userFrom, authResp)
// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
// }

// func TestSendBadAmountOfCoins(t *testing.T) {
// 	userFrom := "userFromBadAmount"
// 	userTo := "userToBadAmount"

// 	authResp, err := authenticate(userFrom, testPassword)
// 	assert.NoError(t, err)

// 	_, err = authenticate(userTo, testPassword)
// 	assert.NoError(t, err)

// 	amounts := []int{-100, 0, 100000}

// 	for _, amount := range amounts {
// 		resp, err := sendCoins(amount, userTo, authResp)
// 		assert.NoError(t, err)
// 		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
// 	}
// }
