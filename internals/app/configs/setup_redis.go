package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func InitializeRedisClient() (*redis.Client, error) {
	log.Println("Setting up Redis client...")
	// Create a new Redis client with the specified options
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PWD"), 
		DB:       0,
	})

	// Ping the Redis server to check if the connection is successful
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		// If there's an error, close the client and return the error
		client.Close()
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	// Return the initialized client and nil error if connection is successful
	return client, nil
}