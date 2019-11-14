package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFor(t *testing.T) {
	i := 1
	for i <= 3 {
		i = i + 1
	}
	assert.Equal(t, 4, i)

	j := 7
	for j = 7; j <= 9; j++ {

	}
	assert.Equal(t, 10, j)

	k := "original"
	for {
		break
		k = "sequel"
	}

	assert.Equal(t, "original", k)

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		assert.Equal(t, 1, n%2)
	}
}
