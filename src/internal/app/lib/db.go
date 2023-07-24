package lib

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Db *sql.DB
}

// NewDatabase creates a new database. The database is connected to the database specified by host user password and dbname
//
// Args:
//
//	host: host to connect to e. g. localhost
//	port: port to connect to e. g. 6379
//
// Returns:
//
//	Database for the new database or nil if there was an error connecting to the database or connection could not be
func NewDatabase(host, user, password, dbname string, port int) Database {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	// if the database connection is not established
	if err != nil {
		log.Panic("error connecting to database")
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	return Database{
		Db: db,
	}
}
