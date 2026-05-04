package cron

import (
	"context"
	"fmt"
	"time"
	"url-shortner/db"
)

func StartCleanupWorker(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := db.DeleteExpired(context.Background())
				if err != nil {
					fmt.Println("Cleanup Error:", err)
				} else {
					fmt.Println("Expired URLs have been cleaned")
				}
			case <-ctx.Done():
				fmt.Println("Cleanup Worker work Done. Stopping Now")
				return
			}
		}
	}()
}
