package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	logs := []string{
		`{"timestamp":"2025-06-20T12:00:00Z","level":"info","msg":"Server started"}`,
		`{"timestamp":"2025-06-20T12:01:00Z","level":"error","msg":"Database timeout"}`,
		`{"timestamp":"2025-06-20T12:02:00Z","level":"info","msg":"Health check passed"}`,
	}

	for _, logEntry := range logs {
		err := PushToQueue(ctx, logEntry)
		if err != nil {
			log.Printf("❌ Failed to queue log: %v", err)
			// In production, you'd retry or store it
		} else {
			fmt.Println("✅ Queued:", logEntry)
		}
	}
}
