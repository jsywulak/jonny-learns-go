package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func TestRecursion(t *testing.T) {
	assert.Equal(t, 5040, fact(7))
}
