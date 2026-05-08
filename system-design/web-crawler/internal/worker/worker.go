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
	"web-crawler/internal/storage/postgres"
	"web-crawler/internal/util"
)

type Worker struct {
	Queue      frontier.Frontier
	Visited    storage.VisitedStore
	SeedDomain string
	Repo       *postgres.URLRepository
	MaxDepth   int
}

func (w *Worker) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := w.Queue.PopContext(ctx)
		if !ok {
			return
		}

		url := task.URL
		normalized, err := util.Normalize(url)
		if err != nil {
			continue
		}

		if w.Visited.Seen(normalized) {
			continue
		}

		if !w.Visited.MarkInProgress(normalized) {
			continue
		}

		exists, err := w.Repo.Exists(normalized)
		if err != nil {
			fmt.Println("exists check failed:", err)
			w.Visited.RemoveInProgress(normalized)
			continue
		}
		if exists {
			w.Visited.RemoveInProgress(normalized)
			continue
		}

		err = w.Repo.InsertURL(normalized, 0)
		if err != nil {
			fmt.Println("insert failed:", err)
			w.Visited.RemoveInProgress(normalized)
			continue
		}

		fmt.Println("Crawling", normalized)

		html, err := fetcher.Fetch(ctx, normalized)
		if err != nil {
			fmt.Println("fetch failed:", err)
			w.Visited.RemoveInProgress(normalized)
			continue
		}

		err = w.Repo.MarkCrawled(normalized)
		if err != nil {
			fmt.Println("mark failed:", err)
		}

		w.Visited.RemoveInProgress(normalized)

		if task.Depth >= w.MaxDepth {
			continue
		}

		links := parser.ExtractLinks(normalized, html)

		for _, link := range links {
			if !filter.SameDomain(w.SeedDomain, link) {
				continue
			}
			if err := w.Queue.Push(frontier.CrawlTask{
				URL:   link,
				Depth: task.Depth + 1,
			}); err != nil {
				fmt.Println("queue push failed:", err)
			}
		}
	}
}
