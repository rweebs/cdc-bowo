package lib

import (
	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Cache *redis.Client
}

func NewCache() Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:36379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return Cache{
		Cache: rdb,
	}
}
