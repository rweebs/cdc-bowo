package utils

import (
	"database/sql"
)

func SQLExecutor(db *sql.DB, statement string) (sql.Result, error) {
	// log.Println(statement)
	result, err := db.Exec(statement)
	return result, err
}
