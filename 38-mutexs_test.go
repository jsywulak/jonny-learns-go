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
	fmt.Println(state, mutex)
}
