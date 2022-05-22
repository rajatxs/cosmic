package storage

import (
	"bytes"
	"os"
	"path"
	"strings"

	"github.com/rajatxs/cosmic/logger"
)

const SqlQueryRootDir = "sql"

func CheckTableExistence(tableName string) bool {
	var count uint8

	row := Sql.QueryRow(
		"SELECT COUNT(name) FROM `sqlite_master` WHERE type=? AND name=?;",
		"table", tableName,
	)

	scanError := row.Scan(&count)

	if scanError != nil {
		logger.Err(scanError)
	}

	return count == 1
}

func getQueryPath(filename *string) string {
	wd, _ := os.Getwd()

	return path.Join(wd, SqlQueryRootDir, *filename+".sql")
}

func ReadQuery(filename string, query *string) error {
	var qpath string = getQueryPath(&filename)
	content, readError := os.ReadFile(qpath)

	if readError != nil {
		logger.Err(readError)
	}

	*query = strings.TrimSpace(bytes.NewBuffer(content).String())

	return readError
}

func ExecQuery(filename string) error {
	var query string

	readError := ReadQuery(filename, &query)

	if readError != nil {
		return readError
	}

	_, execError := Sql.Exec(query)
	return execError
}
