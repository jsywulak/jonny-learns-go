package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeOverChannels(t *testing.T) {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		assert.NotNil(t, elem, "elem should have a value")
		assert.True(t, elem == "one" || elem == "two", "elem should be 'one' or 'two'")
	}
}
