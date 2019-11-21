package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelSync(t *testing.T) {
	done := make(chan bool, 1)
	go worker(done)
	assert.True(t, <-done)
}

func worker(done chan bool) {
	done <- true
}
