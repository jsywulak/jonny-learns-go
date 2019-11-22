package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func waitgroupworker(id int, wg *sync.WaitGroup) {
	wg.Done()
}

func TestWaitGroups(t *testing.T) {
	assert.Equal(t, 1, 1)

	var wg sync.WaitGroup
	workercount := 0
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		workercount++
		go waitgroupworker(i, &wg)
	}
	wg.Wait()
	assert.Equal(t, 5, workercount, "should have five workers")
}
