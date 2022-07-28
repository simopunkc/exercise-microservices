package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	signatureKey []byte = []byte(os.Getenv("JWT_PRIVATE_KEY"))
)

func generateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iss":     "edspert",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(signatureKey)
	if err != nil {
		return "", err
	}
	return stringToken, nil
}
