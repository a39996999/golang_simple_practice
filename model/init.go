package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

type conf struct {
	Host     string
	Database string
	Username string
	Password string
}

var db *sql.DB

func Init() {
	file, err := ioutil.ReadFile("model/conf.json")

	if err != nil {
		panic(err)
	}
	conf := conf{}
	err = json.Unmarshal(file, &conf)
	connection_str := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", conf.Username, conf.Password, conf.Host, conf.Database)
	db, err = sql.Open("mysql", connection_str)

	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
