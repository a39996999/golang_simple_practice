package model

import (
	"chatroom/utils"
	"database/sql"
	"errors"
)

type User struct {
	Id       int
	Name     string
	Password string
	Email    string
	IsVerify bool
	Token    string
}

func CreateUser(username, password, email string) (User, error) {
	searchsql := "select username, email from users where username = ? or email = ?"
	user := User{}
	err := db.QueryRow(searchsql, username, email).Scan(&user.Name, &user.Email)
	if err != sql.ErrNoRows {
		return user, nil
	}
	if err != nil && err != sql.ErrNoRows {
		return user, err
	}
	createTime := utils.GetCurrentTime()
	token, generate_err := utils.GenerateToken()
	if generate_err != nil {
		return user, generate_err
	}
	passwordHash := utils.HashPassword(password, token)
	insertSql := "insert into users(username, password, email, create_time, token) values(?, ?, ?, ?, ?)"
	_, err = db.Exec(insertSql, username, passwordHash, email, createTime, token)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUserPassword(username, password string) error {
	searchsql := "select username from users where username = ?"
	err := db.QueryRow(searchsql, username).Scan(username)
	if err == sql.ErrNoRows {
		return errors.New("user is not exist")
	}
	updatesql := "update users set password = ?, token = ? where username = ?"
	token, generate_err := utils.GenerateToken()
	if generate_err != nil {
		return generate_err
	}
	passwordHash := utils.HashPassword(password, token)
	_, err = db.Exec(updatesql, passwordHash, token, username)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(username string) error {
	deletesql := "delete from users where username = ?"
	searchsql := "select username from users where username = ?"
	err := db.QueryRow(searchsql, username).Scan(username)
	if err == sql.ErrNoRows {
		return errors.New("user is not exist")
	}
	_, err = db.Exec(deletesql, username)
	if err != nil {
		return err
	}
	return nil
}

func QueryUserInfo(username string) (User, error) {
	querySql := "select id, username, email, is_verify_email, password, token from users where username = ?"
	user := User{}
	err := db.QueryRow(querySql, username).Scan(&user.Id, &user.Name, &user.Email, &user.IsVerify, &user.Password, &user.Token)
	if err == sql.ErrNoRows {
		return user, nil
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

func RecordSendMail(email, token string) (bool, error) {
	querySql := "select id, is_verify_email from users where email = ?"
	var user_id int
	var is_verify_email bool
	err := db.QueryRow(querySql, email).Scan(&user_id, &is_verify_email)
	if err != nil {
		return is_verify_email, err
	}
	if is_verify_email != false {
		return is_verify_email, err
	}
	insertSql := "insert into mail(user_id, email, verification_token, create_time) values(?, ?, ?, ?)"
	create_time := utils.GetCurrentTime()
	_, err = db.Exec(insertSql, user_id, email, token, create_time)
	if err != nil {
		return is_verify_email, err
	}
	return is_verify_email, err
}

func VerifyMail(token string) error {
	querySql := "select user_id from mail where verification_token = ?"
	var user_id int
	err := db.QueryRow(querySql, token).Scan(&user_id)
	if err != nil {
		return errors.New("token is not exist")
	}
	updateSql := `update mail
		inner join users on mail.user_id = users.id
		set users.is_verify_email = 1, mail.is_verify = 1
		where mail.verification_token = ?
	`
	_, err = db.Exec(updateSql, token)
	if err != nil {
		return err
	}
	return nil
}
