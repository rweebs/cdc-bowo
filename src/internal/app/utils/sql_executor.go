package utils

import (
	"database/sql"
)

// Executes a SQL statement and returns the result. This is a wrapper around sql. DB. Exec.
//
// Args:
//
//	db: The database to execute the statement on. This must be a pointer to a sql. DB
//	statement: The statement to execute
func SQLExecutor(db *sql.DB, statement string) (sql.Result, error) {
	// log.Println(statement)
	result, err := db.Exec(statement)
	return result, err
}
