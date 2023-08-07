package model

import (
	"chatroom/utils"
	"database/sql"
)

func CreateUser(username, password, email string) error {
	createTime := utils.GetCurrentTime()
	salt, generate_err := utils.GenerateSalt()
	if generate_err != nil {
		return generate_err
	}
	passwordHash := utils.HashPassword(password, salt)
	insertSql := "insert into users(username, password, email, create_time, salt) values(?, ?, ?, ?, ?)"
	_, err := db.Exec(insertSql, username, passwordHash, email, createTime, salt)
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
	if err != nil {
		return err
	}
	return nil
}

func QueryPassword(username string) (string, string, error) {
	querySql := "select password, salt from users where username = ?"
	var password, salt string
	err := db.QueryRow(querySql, username).Scan(&password, &salt)
	if err != nil {
		return "", "", err
	}
	return password, salt, err
}
