package models

import (
	"github.com/go-redis/redis"
)

var db *redis.Client

func Init() {
	db = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Password: "demo", // no password set
		// DB:       0,      // use default DB
	})
}
