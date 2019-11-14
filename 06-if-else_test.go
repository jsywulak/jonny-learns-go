package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfElse(t *testing.T) {
	if 7%2 == 0 {
		assert.Equal(t, 0, 7%2)
	} else {
		assert.Equal(t, 1, 7%2)
	}

	if 8%4 == 0 {
		assert.Equal(t, 0, 8%4)
	}

	if num := 9; num < 10 {
		assert.Equal(t, true, num < 10, "number should be less than 10")
	} else if num < 10 {
		assert.Equal(t, true, num == 0)
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
