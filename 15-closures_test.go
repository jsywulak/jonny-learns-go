package main

import (
	"fmt"
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
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())

}
