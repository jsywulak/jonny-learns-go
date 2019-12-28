package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimits(t *testing.T) {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(20 * time.Millisecond)

	lastTime := time.Now()
	for range requests {
		<-limiter
		now := time.Now()
		assert.WithinDuration(t, lastTime, now, 30*time.Millisecond)
		lastTime = now
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// this is where you planted the tree
	go func() {
		for tt := range time.Tick(20 * time.Millisecond) {
			burstyLimiter <- tt
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	lastTime = time.Now()

	count := 0
	for range burstyRequests {
		<-burstyLimiter
		now := time.Now()
		if count < 3 {
			assert.WithinDuration(t, lastTime, now, 5*time.Millisecond)
		} else {
			assert.WithinDuration(t, lastTime, now, 35*time.Millisecond)
		}
		count++
		lastTime = now
	}
}
