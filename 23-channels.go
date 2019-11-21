package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannels(t *testing.T) {
	messages := make(chan string)

	go func() { messages <- "ping" }()
	msg := <-messages
	assert.Equal(t, "ping", msg)
}
