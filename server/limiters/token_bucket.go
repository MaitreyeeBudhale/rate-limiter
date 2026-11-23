package limiters

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mu sync.Mutex //every goroutine accessing the same TokenBucket instance will use the same mutex, assuming they call the methods on that instance.

	capacity     float64
	tokens       float64
	refillRate   float64
	refillPeriod time.Duration
	lastRefill   time.Time
}

// constructor to create bucket
func NewTokenBucket(capacity int, refillRate float64, refillPeriod time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:     float64(capacity),
		tokens:       float64(capacity),
		refillRate:   refillRate,
		refillPeriod: refillPeriod,
		lastRefill:   time.Now(),
	}
}

// method belonging to struct TokenBucket
// calculate how many tokens are generated within intervals of two successive requests
func (tb *TokenBucket) refill() {
	now := time.Now()

	elapsed := now.Sub(tb.lastRefill)

	tokensToAdd := elapsed.Seconds() /
		tb.refillPeriod.Seconds() * tb.refillRate

	tb.tokens += tokensToAdd

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	tb.lastRefill = now
}

// method belonging to struct TokenBucket
func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()         //acquire lock on bucket
	defer tb.mu.Unlock() //release lock when Allow() finishes

	tb.refill()

	if tb.tokens >= 1 {
		// consume one token
		tb.tokens--
		return true
	}

	return false
}

// total number of request
// bucket capacity
// infow rate
