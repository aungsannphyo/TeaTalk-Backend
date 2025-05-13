package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "this-is-your-secret-key"

func GenerateToken(email string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (string, error) {
	parseToken, error := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if error != nil {
		return "", errors.New("Could not parse token.")
	}

	tokenIsValid := parseToken.Valid
	if !tokenIsValid {
		return "", errors.New("Invalid token!")
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("Invalid token claims.")
	}

	userId := claims["userId"].(string)

	return userId, nil
}
