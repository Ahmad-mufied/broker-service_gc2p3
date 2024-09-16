package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTCustomClaims struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string, username string) (string, error) {
	claims := &JWTCustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as a response
	signedString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return signedString, nil
}
