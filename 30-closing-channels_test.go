package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosingChannels(t *testing.T) {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		count := 0
		previous := 0
		for {
			j, more := <-jobs
			if more {
				assert.True(t, 0 < j && j < 4, "j should be 1-3")
				assert.True(t, j > previous, "j should be consecutive numbers")
				count++
			} else {
				assert.Equal(t, 3, count, "should have received three jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
	}

	close(jobs)
	assert.True(t, <-done, "should have received all jobs")

}
