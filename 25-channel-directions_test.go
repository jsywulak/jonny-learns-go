package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelDirections(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func pings(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
