package main

import (
	"fmt"
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
	assert.True(t, true)
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}
