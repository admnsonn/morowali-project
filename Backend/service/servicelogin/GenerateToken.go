package servicelogin

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateToken(userID int) (string, error) {
	privateKey := []byte("secretKey")

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   fmt.Sprintf("%d", userID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
