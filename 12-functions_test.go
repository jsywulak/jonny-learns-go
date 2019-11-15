package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
func TestFunctions(t *testing.T) {
	assert.Equal(t, 3, plus(1, 2))

	assert.Equal(t, 6, plusPlus(1, 2, 3))

}
