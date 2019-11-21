package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelDirections(t *testing.T) {
	assert.Equal(t, 1, 1)
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
