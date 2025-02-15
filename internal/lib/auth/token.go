package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/qreaqtor/api-avito-shop/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (tm *TokenManager) CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}

func (tm *TokenManager) GenerateToken(username string) (*models.Token, error) {
	iat := time.Now()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iat":      iat.Unix(),
		"exp":      iat.Add(tm.cfg.TokenLifespan).Unix(),
	})

	authToken, err := jwtToken.SignedString([]byte(tm.cfg.SigningKey))
	if err != nil {
		return nil, err
	}

	token := &models.Token{
		AuthToken: authToken,
	}
	return token, nil
}

func (tm *TokenManager) GetHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), tm.cfg.PasswordCostBcrypt)
	return string(hashedPassword), err
}
