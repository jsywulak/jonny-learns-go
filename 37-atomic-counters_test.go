package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicCounters(t *testing.T) {
	assert.Equal(t, 1, 1)
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUnit64(&ops, 1)
			}
			wg.Done()
		}()
	}
	fmt.Print(ops, wg)
}
