package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMutexes(t *testing.T) {
	assert.Equal(t, 1, 1)
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	fmt.Println(state, mutex, readOps, writeOps)
}
