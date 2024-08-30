package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = []byte("secretkey")

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = username
	claims["exp"] = int64(time.Now().Add(time.Hour * 24).Unix())
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return string(tokenString), nil

}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
