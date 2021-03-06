package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func TestVariadicFunctions(t *testing.T) {
	assert.Equal(t, 3, sum(1, 2))
	assert.Equal(t, 6, sum(1, 2, 3))
	nums := []int{1, 2, 3, 4}
	assert.Equal(t, 10, sum(nums...))

}
