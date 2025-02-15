package auth

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const (
	algorithm = "HS256"
	headerPrefix = "Bearer"
)

func (tm *TokenManager) GetContextWithUsername(ctx context.Context, headerValue string) (context.Context, error) {
	tokenStr, err := getToken(strings.Split(headerValue, " "))
	if err != nil {
		return ctx, err
	}

	payload, err := parse(tokenStr, tm.hashSecretGetter)
	if err != nil {
		return ctx, err
	}

	username, err := payload.GetSubject()
	if err != nil {
		return ctx, errBadPayload
	}

	return context.WithValue(ctx, usernameCtxKey, username), nil
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

func getToken(tokenHeader []string) (string, error) {
	if len(tokenHeader) != 2 || tokenHeader[0] != headerPrefix {
		return "", errBadToken
	}
	return tokenHeader[1], nil
}

func (tm *TokenManager) hashSecretGetter(token *jwt.Token) (interface{}, error) {
	method, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok || method.Alg() != algorithm {
		return nil, errBadToken
	}
	return tm.cfg.SigningKey, nil
}
