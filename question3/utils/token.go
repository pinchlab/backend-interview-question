package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtSecret = []byte("secret")

type TokenClaims struct {
	AccountId int `json:"account_id"`
	jwt.StandardClaims
}

func GenerateToken(accountId int) (string, error) {
	now := time.Now()

	tokenClaims := TokenClaims{
		AccountId: accountId,
	}

	tokenClaims.IssuedAt = now.Unix()
	tokenClaims.ExpiresAt = now.Add(time.Hour * 24).Unix()

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	accessToken, err := aToken.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ParseToken(token string) (*TokenClaims, error) {
	tokenClaims := &TokenClaims{}
	_, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return tokenClaims, nil
}
