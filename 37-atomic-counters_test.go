package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicCounters(t *testing.T) {
	assert.Equal(t, 1, 1)
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
	fmt.Println("ops:", opsSafe)
	assert.Equal(t, 50000, opsSafe)
	fmt.Println("ops:", opsDanger)
}
