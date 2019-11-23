package main

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicCounters(t *testing.T) {
	var opsSafe uint64
	var opsDanger uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&opsSafe, 1)
				opsDanger++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(t, uint64(50000), opsSafe, "should've hit 50k safe ops")
	assert.NotEqual(t, uint64(50000), opsDanger, "shouldn't have hit 50k danger ops")
}
