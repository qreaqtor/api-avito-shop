//go:build integration

package integrationstest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/qreaqtor/api-avito-shop/pkg/httprocess"
	"github.com/stretchr/testify/assert"
)

const (
	baseURL = "http://localhost:8080/api"

	testPassword = "passwordtest"
)

type authResponse struct {
	Token string `json:"token"`
}

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	errBadCredentials = errors.New("bad credentials")
	errNoToken        = errors.New("no token in response")
)

func authenticate(username, password string) (*authResponse, error) {
	url := baseURL + "/auth"

	authRequest := &authRequest{
		Username: username,
		Password: password,
	}
	requestBody, err := json.Marshal(authRequest)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, httprocess.ContentTypeJSON, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	authResponse := new(authResponse)
	err = json.NewDecoder(resp.Body).Decode(authResponse)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errBadCredentials
	}
	if authResponse.Token == "" {
		return nil, errNoToken
	}

	return authResponse, nil
}

func TestAuthSimple(t *testing.T) {
	_, err := authenticate("testUsernameSimple", "testPasswordSimple")
	assert.NoError(t, err)
}
