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
		// fmt.Println("no message received")
	}

	msg := "hi"

	select {
	case messages <- msg:
		assert.Fail(t, "should not have gotten into this branch")
		fmt.Println("sent message")
	default:
		assert.True(t, true, "Should have gotten into this branch")
	}

	select {
	case msg := <-messages:
		assert.Fail(t, "should not have gotten into this branch")
		fmt.Println("received message", msg)
	case sig := <-signals:
		assert.Fail(t, "should not have gotten into this branch")
		fmt.Println("recived signal", sig)
	default:
		fmt.Println("no activity")
	}

}
