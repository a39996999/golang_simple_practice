package migrate

import "database/sql"

func CreateUserTable(db *sql.DB) {
	createTableSql := `create table if not exists users(
		id int auto_increment primary key, 
		username varchar(32),
		password varchar(64),
		email varchar(48),
		is_verify_email boolean default false,
		token varchar(32), 
		create_time datetime
	)`
	_, err := db.Exec(createTableSql)
	if err != nil {
		panic(err)
	}
}

func CreateMailTable(db *sql.DB) {
	createTablesql := `create table if not exists mail(
		id int auto_increment primary key,
		user_id int,
		email varchar(48),
		verification_token varchar(32),
		is_verify boolean default false,
		create_time datetime,
		foreign key (user_id) references users(id)
	)`
	_, err := db.Exec(createTablesql)
	if err != nil {
		panic(err)
	}
}

func CreateRoomTable(db *sql.DB) {
	createTableSql := `create table if not exists room(
		id int auto_increment primary key,
		name varchar(32),
		owner_id int,
		description text,
		create_time datetime,
		foreign key (owner_id) references users(id)
	)`
	_, err := db.Exec(createTableSql)
	if err != nil {
		panic(err)
	}
}
