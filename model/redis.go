package model

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var RedisClient *redis.Client

func InitRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "6379"
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	fmt.Println("Connected to Redis successfully")
}
