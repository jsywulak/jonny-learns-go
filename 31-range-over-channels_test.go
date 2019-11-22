package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeOverChannels(t *testing.T) {
	assert.Equal(t, 1, 1)

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}
