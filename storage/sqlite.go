package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Sql *sql.DB

func init() {
	conn, openError := sql.Open("sqlite3", "./app.db")

	if openError != nil {
		panic(openError)
	}

	Sql = conn
}
