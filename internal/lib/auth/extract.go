package auth

import "context"

func ExtractUsername(ctx context.Context) (string, error) {
	username, ok := ctx.Value(usernameCtxKey).(string)
	if !ok {
		return "", errNoUserInContext
	}
	return username, nil
}
