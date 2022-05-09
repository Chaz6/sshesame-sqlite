package main

import (
	"time"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

type LoginData struct {
	Time	time.Time
	Source	string
	User	string
	Pass	string
	Version	string
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}

func CreateTableLogins(db *sql.DB) {
	sql_table := `
	CREATE TABLE IF NOT EXISTS logins(
		id INTEGER PRIMARY KEY,
		time DATETIME,
		source TEXT,
		user TEXT,
		pass TEXT,
		version TEXT
	);
	`

	_, err := db.Exec(sql_table)
	if err != nil { panic(err) }
}

func StoreLogin(db *sql.DB, Source string, User string, Pass string, Version string) {
	if db == nil {
		return;
	}

	sql_additem := `
	INSERT INTO logins(
		time,
		source,
		user,
		pass,
		version
	) values(CURRENT_TIMESTAMP, ?, ?, ?, ?);
	`

	stmt, err := db.Prepare(sql_additem)
	if err != nil { panic(err) }
	defer stmt.Close()

	_, err2 := stmt.Exec(
    Source,
    User,
    Pass,
    Version,
  )
	if err2 != nil { panic(err2) }
}

