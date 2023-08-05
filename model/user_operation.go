package model

import (
	"chatroom/utils"
	"database/sql"
)

func CreateUser(username, password, email string) error {
	createTime := utils.GetCurrentTime()
	insertSql := "insert into users(username, password, email, create_time) values(?, ?, ?, ?)"
	_, err := db.Exec(insertSql, username, password, email, createTime)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserPassword(username, password string) error {
	updatesql := "update users set password = ? where username = ?"
	_, err := db.Exec(updatesql, password, username)
	if err != nil {
		return err
	}
	return nil
}

func FindUserExist(username string) bool {
	searchsql := "select username from users where username = ?"
	err := db.QueryRow(searchsql, username).Scan(username)
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

func DeleteUser(username string) error {
	deletesql := "delete from users where username = ?"
	_, err := db.Exec(deletesql, username)
	if err == nil {
		return err
	}
	return nil
}
