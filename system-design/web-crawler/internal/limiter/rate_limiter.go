package limiter

import (
	"sync"
	"time"
)

type Limiter struct {
	lastAccess map[string]time.Time
	mu         sync.Mutex
	delay      time.Duration
}

func New(delay time.Duration) *Limiter {
	return &Limiter{
		lastAccess: make(map[string]time.Time),
		delay:      delay,
	}
}

func (l *Limiter) Wait(domain string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	last, ok := l.lastAccess[domain]

	if ok {
		sleep := l.delay - time.Since(last)

		if sleep > 0 {
			time.Sleep(sleep)
		}
	}
	l.lastAccess[domain] = time.Now().UTC()
}
