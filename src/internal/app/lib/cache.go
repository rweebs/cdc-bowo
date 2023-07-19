package lib

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Cache *redis.Client
}

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
