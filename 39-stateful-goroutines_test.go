package main

import (
	"fmt"
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

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOps)

	fmt.Println(readOps, writeOps, reads, writes)

}
