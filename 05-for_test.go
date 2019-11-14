package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFor(t *testing.T) {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	assert.Equal(t, 4, i)

	j := 7
	for j = 7; j <= 9; j++ {

	}
	assert.Equal(t, 9, j)
	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
