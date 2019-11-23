package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicCounters(t *testing.T) {
	assert.Equal(t, 1, 1)
	var ops uint64
	var wg sync.Waitgroup
	fmt.Print(ops, wg)
}
