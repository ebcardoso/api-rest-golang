package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtPayload struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

func EncodeJWT(userId string, jwtKey string) (string, error) {
	payload := JwtPayload{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenOutput, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenOutput, err
}

func DecodeJWT(tokenStr string, jwtKey string) (*jwt.Token, JwtPayload, error) {
	payload := JwtPayload{}
	token, err := jwt.ParseWithClaims(tokenStr, &payload, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, JwtPayload{}, err
	}
	return token, payload, nil
}

func VerifyTokenJWT(tokenStr string, jwtKey string) bool {
	token, _, err := DecodeJWT(tokenStr, jwtKey)
	if err != nil || !token.Valid {
		return false
	}
	return true
}
