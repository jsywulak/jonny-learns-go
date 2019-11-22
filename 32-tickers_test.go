package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTickers(t *testing.T) {
	assert.Equal(t, 1, 1)

	ticker := time.NewTicker(10 * time.Millisecond)
	done := make(chan bool)

	go func() {
		count := 0
		for {
			select {
			case <-done:
				assert.Equal(t, 3, count, "should have ticked three times")
				return
			case tick := <-ticker.C:
				assert.NotNil(t, tick, "tick should be valued")
				count++
			}
		}
	}()

	time.Sleep(30 * time.Millisecond)
	ticker.Stop()
	done <- true
}
