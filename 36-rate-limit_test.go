package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimits(t *testing.T) {
	assert.Equal(t, 1, 1)

	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	last_time := time.Now()
	for req := range requests {
		<-limiter
		now := time.Now()
		fmt.Println("request", req, now)
		fmt.Println("diff", now.Sub(last_time))
		last_time = now
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// this is where you planted the tree
	go func() {
		for tt := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- tt
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
