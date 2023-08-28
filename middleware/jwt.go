package middleware

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTVerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token unauthorized",
			})
			c.Abort()
			return
		}
		_, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
	}
}
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		jwt_key := []byte(os.Getenv("jwt_key"))
		return jwt_key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, check := token.Claims.(jwt.MapClaims); check && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Invalid token")
	}
}

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
