package storage

import "sync"

type Visited struct {
	mu   sync.Mutex
	data map[string]struct{}
}

func NewVisited() *Visited {
	return &Visited{
		data: make(map[string]struct{}),
	}
}

func (v *Visited) Seen(url string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()

	if _, ok := v.data[url]; ok {
		return true
	}

	v.data[url] = struct{}{}
	return false
}

func (v *Visited) MarkInProgress(url string) bool {
	v.mu.Lock()
	defer v.mu.Unlock()

	if _, ok := v.data[url]; ok {
		return false
	}
	return true
}

func (v *Visited) RemoveInProgress(url string) error {
	return nil
}
