package cron

import (
	"context"
	"fmt"
	"time"
	"web-crawler/internal/frontier"
	"web-crawler/internal/storage/postgres"
	redisstore "web-crawler/internal/storage/redis"
)

func StartCleanupWorker(
	ctx context.Context,
	interval time.Duration,
	repo *postgres.URLRepository,
	queue *frontier.RedisQueue,
	visited *redisstore.VisitedStore,
) {
	ticker := time.NewTicker(interval)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("Cleaning up Postgres...")
				if err := repo.Truncate(); err != nil {
					fmt.Println("postgres cleanup failed:", err)
				}
				fmt.Println("Cleaning up Redis queue...")
				if err := queue.Cleanup(); err != nil {
					fmt.Println("redis queue cleanup failed:", err)
				}
				fmt.Println("Cleaning up Redis visited set...")
				if err := visited.Cleanup(); err != nil {
					fmt.Println("redis visited cleanup failed:", err)
				}
				fmt.Println("Cleanup complete")
			case <-ctx.Done():
				fmt.Println("Cleanup worker stopped")
				return
			}
		}
	}()
}
