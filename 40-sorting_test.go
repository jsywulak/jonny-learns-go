package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	assert.Equal(t, 1, 1)

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)

}
