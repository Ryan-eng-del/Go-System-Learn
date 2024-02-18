package redis_client

import (
	"time"

	"github.com/go-redis/redis/v8"
)

var Client  *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6380",
		Username: "default",
		Password: "123456",
		DB: 0,
		DialTimeout: 2 * time.Second,
	})

}