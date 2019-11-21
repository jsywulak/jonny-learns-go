package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	assert.Equal(t, 1, 1)
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
}
