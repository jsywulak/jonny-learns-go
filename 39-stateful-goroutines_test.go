package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ReadOp struct {
	key  int
	resp chan int
}

type WriteOp struct {
	key  int
	val  int
	resp chan bool
}

func TestStatefulGoRoutines(t *testing.T) {
	assert.Equal(t, 1, 1)

}
