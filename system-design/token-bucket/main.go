package main

import (
	"fmt"
	"sync"
	"time"
)

// Struct for New Tocken Bucket
type TokenBucket struct {
	capacity   float64
	tokens     float64
	rate       float64
	lastRefill time.Time
	mu         sync.Mutex
}

// Create a new token Bucket with capacity
func NewTokenBucket(capacity, rate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		rate:       rate,
		lastRefill: time.Now(),
	}

}

// Add tokens
func (tb *TokenBucket) refill() {
	now := time.Now()
	elapsedTime := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsedTime * tb.rate

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
}

// check for Request to Go forward
func (tb *TokenBucket) AllowRequest() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill()

	if tb.tokens >= 1 {
		tb.tokens -= 1
		return true
	}
	return false
}

func main() {
	// fmt.Println("Hello WOrld")

	// Create Bucket
	bucket := NewTokenBucket(5, 1)

	for range 30 {
		if bucket.AllowRequest() {
			fmt.Println("Allowed Request, Token Passed")
		} else {
			fmt.Println("Request DENIED!!!")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
