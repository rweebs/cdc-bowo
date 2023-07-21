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

func NewDatabase(host, user, password, dbname string, port int) Database {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Println(psqlconn)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic("error connecting to database")
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	return Database{
		Db: db,
	}
}
