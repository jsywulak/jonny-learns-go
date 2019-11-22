package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func poolworker(id int, jobs <-chan int, results chan<- int) {

}

func TestWorkerPools(t *testing.T) {
	assert.Equal(t, 1, 1)

}
