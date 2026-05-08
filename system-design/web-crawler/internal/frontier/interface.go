package frontier

import "context"

type Frontier interface {
	Push(task CrawlTask) error
	Pop() CrawlTask
	PopContext(ctx context.Context) (CrawlTask, bool)
}
