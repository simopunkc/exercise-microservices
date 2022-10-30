package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type UtilJwt struct {
}

func NewUtilJwt() *UtilJwt {
	return &UtilJwt{}
}

var (
	signatureKey []byte = []byte(os.Getenv("JWT_PRIVATE_KEY"))
)

func (uj UtilJwt) GenerateJWT(userID int64) (string, error) {
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
