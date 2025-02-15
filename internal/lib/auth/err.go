package auth

import "errors"

var (
	errBadToken   = errors.New("bad token")
	errNoPayload  = errors.New("no payload")
	errBadPayload = errors.New("bad value in payload")

	errNoUserInContext = errors.New("username not exists in context")
)
