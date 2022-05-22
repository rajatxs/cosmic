package storage

import "github.com/rajatxs/cosmic/logger"

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
