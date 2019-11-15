package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func vals() (int, int) {
	return 3, 7
}

func TestMultipleReturnValues(t *testing.T) {
	assert.Equal(t, 1, 1)
	a, b := vals()
	assert.Equal(t, 3, a)
	assert.Equal(t, 7, b)

	_, c := vals()
	assert.Equal(t, 7, c)
}
