package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosingChannels(t *testing.T) {
	assert.Equal(t, 1, 1)

	jobs := make(chan int, 5)

	close(jobs)

}
