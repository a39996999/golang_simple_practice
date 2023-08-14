package utils

import (
	"net/smtp"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type config struct {
	SmtpConfig smtpconfig `yaml:"smtp"`
}

type smtpconfig struct {
	From       string `yaml:"from"`
	Password   string `yaml:"password"`
	SmtpServer string `yaml:"smtpServer"`
	Port       string `yaml:"port"`
}

var conf config
var once sync.Once

func InitMailConfig() {
	once.Do(func() {
		configfile, err := os.Open("config.yml")
		if err != nil {
			panic(err)
		}
		defer configfile.Close()
		decoder := yaml.NewDecoder(configfile)
		err = decoder.Decode(&conf)
		if err != nil {
			panic(err)
		}
	})
}

func SendMail(to string, message []byte) error {
	InitMailConfig()
	auth := smtp.PlainAuth("", conf.SmtpConfig.From, conf.SmtpConfig.Password, conf.SmtpConfig.SmtpServer)
	return smtp.SendMail(conf.SmtpConfig.SmtpServer+":"+conf.SmtpConfig.Port, auth, conf.SmtpConfig.From, []string{to}, message)
}
