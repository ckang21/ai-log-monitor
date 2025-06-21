package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379", // Redis must be running locally for now
	DB:   0,
})

func PushToQueue(ctx context.Context, logLine string) error {
	const maxRetries = 3
	const retryDelay = time.Second * 2

	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Simulate occasional failure only for a specific log line
		if logLine == `{"timestamp":"2025-06-20T12:01:00Z","level":"error","msg":"Database timeout"}` && attempt < maxRetries {
			err = fmt.Errorf("simulated failure on attempt %d", attempt)
		} else {
			err = rdb.RPush(ctx, "log_queue", logLine).Err()
		}

		if err == nil {
			return nil
		}

		fmt.Printf("⚠️  Attempt %d failed: %v\n", attempt, err)
		time.Sleep(retryDelay)
	}

	return fmt.Errorf("failed after %d retries: %w", maxRetries, err)
}
