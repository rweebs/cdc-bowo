package utils

import (
	"database/sql"
	"fmt"
)

func SQLExecutor(db *sql.DB, statement string) (sql.Result, error) {
	fmt.Println(statement)
	result, err := db.Exec(statement)
	return result, err
}
