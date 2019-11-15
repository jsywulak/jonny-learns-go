package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrays(t *testing.T) {
	var a [5]int
	fmt.Println("emp:", a)
	assert.ElementsMatch(t, a, [5]int{0, 0, 0, 0, 0})

	a[4] = 100
	assert.Equal(t, 100, a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dc1:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
