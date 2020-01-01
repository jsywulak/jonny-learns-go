package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var workercount int = 0

func waitgroupworker(id int, wg *sync.WaitGroup) {
	workercount++
	time.Sleep(time.Millisecond)
	wg.Done()
}

func TestWaitGroups(t *testing.T) {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go waitgroupworker(i, &wg)
	}
	wg.Wait()
	time.Sleep(time.Millisecond * 100)
	assert.Equal(t, 5, workercount, "should have five workers")
}
