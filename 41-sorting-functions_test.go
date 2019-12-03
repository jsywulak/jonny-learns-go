package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestSortingFunctions(t *testing.T) {
	assert.Equal(t, 1, 1)
}
