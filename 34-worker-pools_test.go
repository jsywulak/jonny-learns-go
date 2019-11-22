package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func poolworker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		// fmt.Println("worker", id, "started job", j)
		// time.Sleep(time.Second)
		// fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}

}

func TestWorkerPools(t *testing.T) {
	assert.Equal(t, 1, 1)

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go poolworker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println("a[", a, "] is ", <-results)
	}

}
