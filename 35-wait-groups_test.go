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
}
