package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library and do not remove on save
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createTableUser := `
	create table if not exists users (
	id integer primary key autoincrement,
	password text not null,
	email text not null unique
	)`

	_, err := DB.Exec(createTableUser)
	if err != nil {
		panic("could not create user table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create event table")
	}

	createTableRegisterion := `
	create table if not exists registerations (
	id integer primary key autoincrement,
	event_id integer not null,
	user_id integer not null,
	foreign key(event_id) references events(id),
	foreign key(user_id) references users(id)
	)
	`
	_, err = DB.Exec(createTableRegisterion)
	if err != nil {
		panic("could not create registration table")
	}
}
