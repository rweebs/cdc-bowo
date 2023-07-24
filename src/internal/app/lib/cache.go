package lib

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Cache *redis.Client
}

// NewCache creates a new cache. If password is empty it will use default password. This is useful for test cases that don t care about password based caching
//
// Args:
//
//	host: Redis host e. g. localhost
//	port: Redis port e. g. 6379
//	password: Redis password e. g. password for user password
//
// Returns:
//
//	NewCache returns a new cache for the given host and port and password or nil if there is no
func NewCache(host string, port int, password string) Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password, // no password set
		DB:       0,        // use default DB
	})

	return Cache{
		Cache: rdb,
	}
}
