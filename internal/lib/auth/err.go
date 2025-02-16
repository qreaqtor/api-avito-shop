package auth

import "errors"

var (
	errBadToken   = errors.New("bad token")
	errNoPayload  = errors.New("no payload")
	errBadPayload = errors.New("bad value in payload")
	errNotValidToken = errors.New("token not valid")

	errNoUserInContext = errors.New("username not exists in context")
)
