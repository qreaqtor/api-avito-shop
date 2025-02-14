package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/qreaqtor/api-avito-shop/internal/config"
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type TokenManager struct {
	cfg config.AuthConfig
}

func NewTokenManager() *TokenManager {
	return &TokenManager{}
}

func (tm *TokenManager) CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}

func (tm *TokenManager) GenerateToken(username string) (*models.Token, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(tm.cfg.TokenLifespan).Unix()

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	authToken, err := jwtToken.SignedString([]byte(tm.cfg.Secret))
	if err != nil {
		return nil, err
	}

	token := &models.Token{
		AuthToken: authToken,
	}
	return token, nil
}

func (tm *TokenManager) GetHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}
