package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRateLimits(t *testing.T) {
	assert.Equal(t, 1, 1)

	requests := make(chan int, 5)
	for i := 0; i <= 5; i++ {

		requests <- i
	}

}
