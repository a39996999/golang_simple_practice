package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func GenerateToken() (string, error) {
	saltByte := make([]byte, 16)
	_, err := rand.Read(saltByte)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(saltByte), err
}

func HashPassword(password, salt string) string {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(password + salt))
	return fmt.Sprintf("%x", sha256Hash.Sum(nil))
}
