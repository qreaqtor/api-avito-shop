package auth

import (
	"context"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	algorithm    = "HS256"
	headerPrefix = "Bearer"
)

func (tm *TokenManager) GetContextWithUsername(ctx context.Context, headerValue string) (context.Context, error) {
	tokenStr, err := getToken(headerValue)
	if err != nil {
		return nil, err
	}

	payload, err := parse(tokenStr, tm.hashSecretGetter)
	if err != nil {
		return nil, err
	}

	if !isPayloadValid(payload) {
		return nil, errNotValidToken
	}

	username, ok := payload["username"].(string)
	if !ok || username == "" {
		return nil, errBadPayload
	}

	return context.WithValue(ctx, usernameCtxKey, username), nil
}

func isPayloadValid(payload jwt.MapClaims) bool {
	expUnix, ok := payload["exp"].(float64)
	if !ok {
		return false
	}

	iatUnix, ok := payload["iat"].(float64)
	if !ok {
		return false
	}

	now := time.Now()
	exp := time.Unix(int64(expUnix), 0)
	iat := time.Unix(int64(iatUnix), 0)

	return iat.Before(now) && now.Before(exp)
}

func parse(tokenStr string, keyFunc jwt.Keyfunc) (jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errNoPayload
	}
	return payload, nil
}

func getToken(header string) (string, error) {
	headerSplit := strings.Split(header, " ")
	if len(headerSplit) != 2 || headerSplit[0] != headerPrefix {
		return "", errBadToken
	}
	return headerSplit[1], nil
}

func (tm *TokenManager) hashSecretGetter(token *jwt.Token) (interface{}, error) {
	method, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok || method.Alg() != algorithm {
		return nil, errBadToken
	}
	return []byte(tm.cfg.SigningKey), nil
}
