package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariables(t *testing.T) {
	var a = "initial"
	assert.Equal(t, "initial", a)

	var b, c int = 1, 2
	assert.Equal(t, 1, b)
	assert.Equal(t, 2, c)

	var d = true
	assert.Equal(t, true, d)

	var e int
	assert.Equal(t, 0, e)

	f := "apple"
	assert.Equal(t, "apple", f)

	var g = "orange"
	assert.Equal(t, "orange", g)
}
