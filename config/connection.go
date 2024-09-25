package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var RedisConn *redis.Client

var ctx = context.Background()

func CreateConnection() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	RedisConn = rdb
}

func CloseConnection() {
	err := RedisConn.Close()
	if err != nil {
		fmt.Println("Error closing redis connection")
	}
}

func Upsert(key string, value string) {
	err := RedisConn.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println("Error upserting key")
	}
}

func GetValue(key string) (string, error) {
	val, err := RedisConn.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("Error getting key")
		return "", err
	}
	return val, nil
}
