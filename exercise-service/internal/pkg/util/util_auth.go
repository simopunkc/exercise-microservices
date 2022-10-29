package util

import (
	"errors"
	"os"

	jwt "github.com/golang-jwt/jwt/v4"
)

var (
	privateKey []byte = []byte(os.Getenv("JWT_PRIVATE_KEY"))
)

func DecriptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		}
		return privateKey, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}

	if !parsedToken.Valid {
		return data, errors.New("invalid token")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
