package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelBuffering(t *testing.T) {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	assert.Equal(t, "buffered", <-messages)
	assert.Equal(t, "channel", <-messages)
}
