package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChannelSync(t *testing.T) {
	assert.Equal(t, 1, 1)
	done := make(chan bool, 1)
	go worker(done)
	assert.True(t, <-done)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
