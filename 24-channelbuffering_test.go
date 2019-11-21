package main

import (
	"fmt"
	"testing"
)

func TestChannelBuffering(t *testing.T) {
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
