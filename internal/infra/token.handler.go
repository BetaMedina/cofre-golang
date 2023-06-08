package infra

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenInfra interface {
	GenerateToken(userId string) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type tokenInfra struct{}

func (t tokenInfra) GenerateToken(userId string) (string, error) {
	claims :=
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    userId,
		}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetEnvs("SIGNED_STRING")))
}

func (t tokenInfra) VerifyToken(tokenString string) (*jwt.Token, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetEnvs("SIGNED_STRING")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Invalid token has been provided")
	}

	return token, nil
}

func NewTokenInfra() TokenInfra {
	return &tokenInfra{}
}
