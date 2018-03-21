package redis

import (
	"github.com/gomodule/redigo/redis"
)

var (
	Pool *redis.Pool
)

func init() {
	// redisHost := os.Getenv("REDIS_HOST")
	// if redisHost == "" {
	// 	redisHost := ":6379"
	// }
}

// func newPool(serve string) *redis.Pool {

// }
