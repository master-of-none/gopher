package frontier

import "context"

type InMemoryQueue struct {
	ch chan CrawlTask
}

func New(size int) *InMemoryQueue {
	return &InMemoryQueue{
		ch: make(chan CrawlTask, size),
	}
}

func (q *InMemoryQueue) Push(task CrawlTask) error {
	select {
	case q.ch <- task:
	default:
	}
	return nil
}

func (q *InMemoryQueue) Pop() CrawlTask {
	return <-q.ch
}

func (q *InMemoryQueue) PopContext(ctx context.Context) (CrawlTask, bool) {
	select {
	case task := <-q.ch:
		return task, true
	case <-ctx.Done():
		return CrawlTask{}, false
	}
}
