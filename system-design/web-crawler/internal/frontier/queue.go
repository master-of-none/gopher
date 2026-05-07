package frontier

type InMemoryQueue struct {
	ch chan string
}

func New(size int) *InMemoryQueue {
	return &InMemoryQueue{
		ch: make(chan string, size),
	}
}

func (q *InMemoryQueue) Push(url string) {
	select {
	case q.ch <- url:
	default:
	}
}

func (q *InMemoryQueue) Pop() string {
	return <-q.ch
}
