package model

import (
	"chatroom/model/migrate"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

type config struct {
	Sql databaseConfig `yaml:"mysql"`
}

type databaseConfig struct {
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var db *sql.DB

func Init() {
	configFile, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()
	config := config{}
	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	connection_str := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.Sql.Username, config.Sql.Password, config.Sql.Host, config.Sql.Database)
	db, err = sql.Open("mysql", connection_str)
	if err != nil {
		panic(err)
	}
	migrate.CreateUserTable(db)
	migrate.CreateMailTable(db)
}
