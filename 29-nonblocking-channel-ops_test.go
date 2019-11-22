package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonblockingChannelOps(t *testing.T) {
	assert.Equal(t, 1, 1)
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages:
		assert.Fail(t, "should not have gotten into this branch")
		fmt.Println("received message", msg)
	default:
		assert.True(t, true, "Should have gotten into this branch")
		fmt.Println("no message received")
	}

	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("sent message")
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("recived signal", sig)
	default:
		fmt.Println("no activity")
	}

}
