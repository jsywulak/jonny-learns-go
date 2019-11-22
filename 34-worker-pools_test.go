package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func poolworker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
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
		expected := [5]int{2, 4, 6, 8, 10}
		assert.Contains(t, expected, <-results, "result not in expected result set")
	}

}
