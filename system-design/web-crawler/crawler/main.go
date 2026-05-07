package main

import (
	"context"
	"sync"
	"time"
	"web-crawler/internal/frontier"
	"web-crawler/internal/storage"
	"web-crawler/internal/worker"
)

const workerCount = 5

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	queue := frontier.New(1000)
	queue.Push("https://www.vercel.com")

	visited := storage.NewVisited()

	w := &worker.Worker{
		Queue:   queue,
		Visited: visited,
	}

	var wg sync.WaitGroup

	for range workerCount {
		wg.Add(1)
		go w.Start(ctx, &wg)
	}

	time.Sleep(10 * time.Second)

	cancel()

	wg.Wait()
}
