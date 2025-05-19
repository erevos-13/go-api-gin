package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET = "secretmysectrer"

func GenerateJwtToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(SECRET))

}

func ValidateJwtToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(SECRET), nil
	})
	if err != nil {
		return 0, errors.New("invalid token")
	}
	tokenIsValid := token.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims) //INFO is type checking
	if !ok {
		return 0, errors.New("invalid token")
	}
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
