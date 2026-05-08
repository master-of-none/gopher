package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"web-crawler/cron"
	"web-crawler/internal/frontier"
	"web-crawler/internal/storage/postgres"
	redisstore "web-crawler/internal/storage/redis"
	"web-crawler/internal/worker"

	goredis "github.com/redis/go-redis/v9"
)

const workerCount = 5
const maxDepth = 2

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	postgres.InitDB()
	defer postgres.CloseDB()

	rdb := goredis.NewClient(&goredis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	repo := postgres.NewURLRepository(postgres.Pool)
	seed := "https://pkg.go.dev"
	queue := frontier.NewRedisQueue(rdb)
	if err := repo.Truncate(); err != nil {
		panic(err)
	}
	if err := queue.Cleanup(); err != nil {
		panic(err)
	}
	visited := redisstore.NewVisitedStore(rdb)
	if err := visited.Cleanup(); err != nil {
		panic(err)
	}
	if err := queue.Push(frontier.CrawlTask{
		URL:   seed,
		Depth: 0,
	}); err != nil {
		panic(err)
	}

	w := &worker.Worker{
		Queue:      queue,
		Visited:    visited,
		SeedDomain: seed,
		Repo:       repo,
		MaxDepth:   maxDepth,
	}

	var wg sync.WaitGroup

	for range workerCount {
		wg.Add(1)
		go w.Start(ctx, &wg)
	}

	cron.StartCleanupWorker(ctx, 1*time.Minute, repo, queue, visited)
	<-ctx.Done()
	wg.Wait()
}
