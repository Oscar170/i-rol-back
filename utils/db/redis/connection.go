package redis

import (
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func Save(key string, value interface{}) error {
	err := client.Set(key, value, 0).Err()
	return err
}

func Restore(key string) (string, error) {
	val, err := client.Get(key).Result()
	return val, err
}
