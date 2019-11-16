package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestClosures(t *testing.T) {
	assert.Equal(t, 4, 4)
	nextInt := intSeq()
	assert.Equal(t, 1, nextInt())
	assert.Equal(t, 2, nextInt())
	assert.Equal(t, 3, nextInt())
	assert.Equal(t, 4, nextInt())

	newInts := intSeq()
	assert.Equal(t, 1, newInts())
	assert.Equal(t, 2, newInts())

}
