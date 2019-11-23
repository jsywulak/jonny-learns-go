package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		// time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		// time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			assert.Equal(t, "one", msg1)
		case msg2 := <-c2:
			assert.Equal(t, "two", msg2)
		}
	}
}
