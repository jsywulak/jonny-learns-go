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
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	close(jobs)

}
