package middleware

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user_id int, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	jwt_key := []byte(os.Getenv("jwt_key"))
	tokenString, err := token.SignedString(jwt_key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
