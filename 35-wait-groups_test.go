package main

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func waitgroupworker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d complete", id)
	wg.Done()
}

func TestWaitGroups(t *testing.T) {
	assert.Equal(t, 1, 1)
}
