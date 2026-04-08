package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work completed")
	case <-ctx.Done():
		fmt.Println("Work cancelled:", ctx.Err())
	}
}

func main() {
	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	doWork(ctx)

	// Context with value
	ctx2 := context.WithValue(context.Background(), "userID", "12345")
	userID := ctx2.Value("userID")
	fmt.Printf("User ID: %v\n", userID)
}