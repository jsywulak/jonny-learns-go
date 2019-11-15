package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrays(t *testing.T) {
	var a [5]int
	assert.ElementsMatch(t, a, [5]int{0, 0, 0, 0, 0})

	a[4] = 100
	assert.Equal(t, 100, a[4])
	assert.Equal(t, 5, len(a))

	b := [5]int{1, 2, 3, 4, 5}
	assert.ElementsMatch(t, b, [5]int{1, 2, 3, 4, 5})

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	assert.ElementsMatch(t, [2][3]int{{0, 1, 2}, {1, 2, 3}}, twoD)
}
