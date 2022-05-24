package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rajatxs/cosmic/logger"
)

var Sql *sql.DB

func init() {
	conn, openError := sql.Open("sqlite3", "./app.db")
	conn.SetMaxIdleConns(1)

	if openError != nil {
		logger.Err(openError)
	}

	Sql = conn
}
