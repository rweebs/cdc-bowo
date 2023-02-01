package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/rweebs/cdc-bowo/internal/app/utils"
)

func initRedisStreamList(db *sql.DB, rdb *redis.Client, key string) map[string]string {
	val, _ := rdb.Get(context.Background(), key).Result()
	streams := map[string]string{}
	if json.Unmarshal([]byte(val), &streams) == nil {
		fmt.Println("dari redis")
		return streams

	}
	rows, err := db.Query(`select CONCAT(schemaname,'.',tablename) from pg_publication_tables where pubname='dbz_publication';`)
	utils.CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var tableName string

		err = rows.Scan(&tableName)
		utils.CheckError(err)

		streams[fmt.Sprintf("test2.%s", tableName)] = "0"
	}
	return streams
}

func initPrimaryKeyList(db *sql.DB) map[string][]string {
	primaryKeyList := map[string][]string{}
	rows, err := db.Query(`select CONCAT(table_schema,'.',table_name) as table ,column_name FROM information_schema.constraint_column_usage where constraint_name like '%pkey' order by table_schema, table_name;`)
	utils.CheckError(err)
	defer rows.Close()
	for rows.Next() {
		var tableName string
		var columnName string

		err = rows.Scan(&tableName, &columnName)
		utils.CheckError(err)

		val, ok := primaryKeyList[tableName]
		if ok {
			val = append(val, columnName)
			primaryKeyList[tableName] = val
		} else {
			primaryKeyList[tableName] = []string{columnName}
		}
	}
	fmt.Println(primaryKeyList)
	return primaryKeyList
}

func makeStream(streams map[string]string) []string {
	top := []string{}
	bottom := []string{}
	for key, val := range streams {
		top = append(top, key)
		bottom = append(bottom, val)
	}
	return append(top, bottom...)

}
