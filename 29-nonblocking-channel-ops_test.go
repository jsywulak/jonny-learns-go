package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonblockingChannelOps(t *testing.T) {
	assert.Equal(t, 1, 1)
	messages := make(chan string)
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("sent message")
	default:
		fmt.Println("no message sent")
	}

}
