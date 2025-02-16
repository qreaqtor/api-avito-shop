//go:build integration

package integrationstest

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	clientCoins = 1000
)

func buyItem(item string, auth *authResponse) (*http.Response, error) {
	url := fmt.Sprintf("%s/buy/%s", baseURL, item)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", "Bearer "+auth.Token)

	return http.DefaultClient.Do(req)
}

func TestBuyMerchSimple(t *testing.T) {
	username := "userMerchSimple"

	authResp, err := authenticate(username, testPassword)
	assert.NoError(t, err)

	item := "pen"

	buyResp, err := buyItem(item, authResp)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, buyResp.StatusCode)
}
