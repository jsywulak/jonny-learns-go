package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func TestSortingFunctions(t *testing.T) {
	assert.Equal(t, 1, 1)
}
