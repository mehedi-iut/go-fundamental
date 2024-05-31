package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "secret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return err
	}
	tokenIsValid := token.Valid
	if !tokenIsValid {
		return errors.New("Invalid token")
	}

	//claims, ok := token.Claims.(jwt.MapClaims)
	//if !ok {
	//	return nil, errors.New("Invalid token")
	//}
	//
	//email := claims["email"].(string)
	//userId := int64(claims["userId"].(int64))

	return nil
}
