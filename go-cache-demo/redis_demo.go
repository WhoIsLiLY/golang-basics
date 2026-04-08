package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Test connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// Set value
	err = rdb.Set(ctx, "user:1", "John Doe", time.Hour).Err()
	if err != nil {
		fmt.Printf("Error setting value: %v\n", err)
		return
	}

	// Get value
	val, err := rdb.Get(ctx, "user:1").Result()
	if err != nil {
		fmt.Printf("Error getting value: %v\n", err)
		return
	}
	fmt.Printf("Retrieved value: %s\n", val)

	// Delete value
	err = rdb.Del(ctx, "user:1").Err()
	if err != nil {
		fmt.Printf("Error deleting value: %v\n", err)
		return
	}

	fmt.Println("Cache operations completed!")
}