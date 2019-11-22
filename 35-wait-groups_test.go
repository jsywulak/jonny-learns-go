package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var workercount int = 0

func waitgroupworker(id int, wg *sync.WaitGroup) {
	workercount++
	wg.Done()
}

func TestWaitGroups(t *testing.T) {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go waitgroupworker(i, &wg)
	}
	wg.Wait()
	assert.Equal(t, 5, workercount, "should have five workers")
}
