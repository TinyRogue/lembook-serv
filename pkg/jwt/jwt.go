package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var (
	SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
)

const (
	dayInHours = 24
	weekInDays = 7
)

func GenerateToken(username *string) (*string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = *username
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * dayInHours * weekInDays).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func ParseToken(tokenStr string) (*string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["uid"].(string)
		return &username, nil
	} else {
		return nil, err
	}
}
