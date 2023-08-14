package model

import (
	"chatroom/model/migrate"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	connection_str := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("mysql_username"), os.Getenv("mysql_password"), os.Getenv("mysql_host"), os.Getenv("mysql_database"))
	db, err = sql.Open("mysql", connection_str)
	if err != nil {
		panic(err)
	}
	migrate.CreateUserTable(db)
	migrate.CreateMailTable(db)
}
