package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMutexes(t *testing.T) {
	assert.Equal(t, 1, 1)
	var state = make(map[int]int)
	fmt.Println(state)
}
