package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const KEY = `EURFGKHJOUIYJGVHYRGFHBVNJHIJKVNCHDYRHGCBNVJGUYJGNVBC`

func GenerateToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(KEY))
}
func VerifyToken(token string) (any, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New(`Unauthorized Method`)
		}
		return []byte(KEY), nil
	})
	if err != nil {
		fmt.Println("An Error Occured")
		return nil, errors.New("")
	}
	if !parsedToken.Valid {
		fmt.Println("An Error Occured")
		return nil, errors.New("")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("An Error Occured")
		return nil, errors.New("")
	}

	return int64(claims["id"].(float64)), nil
}
