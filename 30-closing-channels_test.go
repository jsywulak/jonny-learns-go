package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosingChannels(t *testing.T) {
	assert.Equal(t, 1, 1)

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		count := 0
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
				count++
			} else {
				fmt.Println("received all jobs")
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
