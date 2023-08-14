package utils

import (
	"net/smtp"
	"os"
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

func SendMail(to string, message []byte) error {
	auth := smtp.PlainAuth("", os.Getenv("smtp_from"), os.Getenv("smtp_password"), os.Getenv("smtp_server"))
	return smtp.SendMail(os.Getenv("smtp_server")+":"+os.Getenv("smtp_port"), auth, os.Getenv("smtp_from"), []string{to}, message)
}
