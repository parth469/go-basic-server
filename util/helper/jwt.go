package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/parth469/go-basic-server/util/config"
)

func CreateToken(data any) (string, error) {
	secretKey := []byte(config.App.SecretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"data": data,
			"exp":  time.Now().Add(5 * time.Minute).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (map[string]interface{}, error) {
	secretKey := []byte(config.App.SecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims format")
	}

	payload, ok := claims["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("payload 'data' not found or invalid in token")
	}

	return payload, nil
}
