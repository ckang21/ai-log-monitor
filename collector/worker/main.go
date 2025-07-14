package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/redis/go-redis/v9"
)

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Msg       string `json:"msg"`
}

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // same as collector
		DB:   0,
	})

	client := resty.New()

	fmt.Println("🔁 Starting log worker...")

	for {
		// Pull one log from Redis (blocking for up to 5 sec)
		result, err := rdb.BLPop(ctx, 5*time.Second, "log_queue").Result()
		if err != nil {
			if err == redis.Nil {
				continue // no log, just wait again
			}
			log.Println("Redis error:", err)
			continue
		}

		logStr := result[1]
		var logEntry LogEntry
		if err := json.Unmarshal([]byte(logStr), &logEntry); err != nil {
			log.Println("Invalid log format:", err)
			continue
		}

		// Send to Python API
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(logEntry).
			SetResult(map[string]bool{}).
			Post("http://localhost:8000/analyze")

		if err != nil {
			log.Println("Failed to call detector:", err)
			continue
		}

		resultMap := resp.Result().(*map[string]bool)
		isAnomaly := (*resultMap)["anomaly"]

		if isAnomaly {
			log.Println("🚨 Anomaly Detected:", logEntry.Msg)
		} else {
			log.Println("✅ Normal log:", logEntry.Msg)
		}
	}
}
