package database

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       dbNo,
	})
	s, err := rdb.Ping(context.Background()).Result()
	fmt.Println("print ping result from database = ", s)
	fmt.Println("Check redis client = ", err)
	fmt.Println("print the created redis client = ", rdb)
	return rdb
}
