package worker

import (
	"context"
	"fmt"
	"sync"
	"web-crawler/internal/fetcher"
	"web-crawler/internal/filter"
	"web-crawler/internal/frontier"
	"web-crawler/internal/parser"
	"web-crawler/internal/storage"
	"web-crawler/internal/util"
)

type Worker struct {
	Queue      *frontier.InMemoryQueue
	Visited    *storage.Visited
	SeedDomain string
}

func (w *Worker) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			url := w.Queue.Pop()

			normalized, err := util.Normalize(url)

			if err != nil {
				continue
			}

			if w.Visited.Seen(normalized) {
				continue
			}
			fmt.Println("Crawling", normalized)

			html, err := fetcher.Fetch(ctx, normalized)

			if err != nil {
				continue
			}

			links := parser.ExtractLinks(normalized, html)

			for _, link := range links {
				if !filter.SameDomain(w.SeedDomain, link) {
					continue
				}
				w.Queue.Push(link)
			}
		}

	}
}
