package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelDirections(t *testing.T) {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	assert.Equal(t, "passed message", <-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
