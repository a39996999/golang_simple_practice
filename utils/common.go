package utils

import (
	"regexp"
	"time"
)

func GetCurrentTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime
}

func VerifyEmailFormat(email string) (bool, error) {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	checkEmailValidate, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false, err
	}
	return checkEmailValidate, nil
}
