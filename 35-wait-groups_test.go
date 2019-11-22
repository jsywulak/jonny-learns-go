package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func waitgroupworker(id int, wg *sync.WaitGroup) {
	// fmt.Printf("Worker %d starting\n", id)
	// time.Sleep(time.Second)
	// fmt.Printf("Worker %d complete\n", id)
	wg.Done()
}

func TestWaitGroups(t *testing.T) {
	assert.Equal(t, 1, 1)

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go waitgroupworker(i, &wg)
	}
	wg.Wait()
}
