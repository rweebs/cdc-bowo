package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/rweebs/cdc-bowo/internal/app/utils"
)

// initRedisStreamList returns a map of stream names to stream ids for a given public name that is stored in Redis
//
// Args:
//
//	db: the database to connect to
//	rdb: the redis connection to use
//	key: the key to read the streams from redis. If it doesn t exist it will create it. This is used to check if we have a connection to the redis server
//	pubName: the name of the publication we are looking for
//	redisPrefix: the prefix to use for the stream ids
func initRedisStreamList(db *sql.DB, rdb *redis.Client, key string, pubName string, redisPrefix string) map[string]string {
	// log.Println(redisPrefix)
	val, _ := rdb.Get(context.Background(), key).Result()
	streams := map[string]string{}
	// Return a list of streams from the redis
	if json.Unmarshal([]byte(val), &streams) == nil {
		// log.Println("dari redis")
		return streams

	}
	rows, err := db.Query(fmt.Sprintf(`select CONCAT(schemaname,'.',tablename) from pg_publication_tables where pubname='%s';`, pubName))
	utils.CheckError(err)
	defer rows.Close()
	// Read the next row from the rows array and store it in streams.
	for rows.Next() {
		var tableName string

		err = rows.Scan(&tableName)
		utils.CheckError(err)

		streams[fmt.Sprintf("%s.%s", redisPrefix, tableName)] = "0"
	}
	return streams
}

// initPrimaryKeyList returns a map of table name to primary key values for all tables in the database.
//
// Args:
//
//	db: the database to query the information_schema. constraint_column_
func initPrimaryKeyList(db *sql.DB) map[string][]string {
	primaryKeyList := map[string][]string{}
	rows, err := db.Query(`select CONCAT(table_schema,'.',table_name) as table ,column_name FROM information_schema.constraint_column_usage where constraint_name like '%pkey' order by table_schema, table_name;`)
	utils.CheckError(err)
	defer rows.Close()
	// Get the primary key list for each row.
	for rows.Next() {
		var tableName string
		var columnName string

		err = rows.Scan(&tableName, &columnName)
		utils.CheckError(err)

		val, ok := primaryKeyList[tableName]
		// Add a column name to primaryKeyList.
		if ok {
			val = append(val, columnName)
			primaryKeyList[tableName] = val
		} else {
			primaryKeyList[tableName] = []string{columnName}
		}
	}

	primaryKeyList["public.t2"] = []string{"id"}
	// log.Println(primaryKeyList)
	return primaryKeyList
}

// Makes a stream from a map of key value pairs. This is used to create an input stream to send to the API
//
// Args:
//
//	streams: map of key value
func makeStream(streams map[string]string) []string {
	top := []string{}
	bottom := []string{}
	// appends the streams to the top and bottom streams
	for key, val := range streams {
		top = append(top, key)
		bottom = append(bottom, val)
	}
	return append(top, bottom...)

}
