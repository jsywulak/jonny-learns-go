package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTickers(t *testing.T) {
	assert.Equal(t, 1, 1)

	ticker := time.NewTicker(10 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case tick := <-ticker.C:
				fmt.Println("Ticker at", tick)
				assert.NotNil(t, tick, "tick should be valued")
			}
		}
	}()

	time.Sleep(30 * time.Millisecond)
	ticker.Stop()
	done <- true
}
