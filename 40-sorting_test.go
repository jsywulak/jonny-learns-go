package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	strs := []string{"c", "a", "b"}
	assert.False(t, sort.StringsAreSorted(strs))

	sort.Strings(strs)

	assert.True(t, sort.StringsAreSorted(strs))
	assert.Equal(t, []string{"a", "b", "c"}, strs)

	ints := []int{7, 2, 4}
	assert.False(t, sort.IntsAreSorted(ints))

	sort.Ints(ints)

	assert.True(t, sort.IntsAreSorted(ints))
	assert.Equal(t, []int{2, 4, 7}, ints)
}
