package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	assert.Equal(t, 1, 1)

	strs := []string{"c", "a", "b"}
	assert.False(t, sort.StringsAreSorted(strs))
	sort.Strings(strs)
	assert.True(t, sort.StringsAreSorted(strs))
	assert.Equal(t, []string{"a", "b", "c"}, strs)

	ints := []int{7, 2, 4}
	assert.False(t, sort.IntsAreSorted(ints))
	is := sort.IntsAreSorted(ints)

	fmt.Println("ints:", ints)
	fmt.Println("sorted?", is)

	sort.Ints(ints)
	assert.True(t, sort.IntsAreSorted(ints))
	assert.Equal(t, []int{2, 4, 7}, ints)
	is = sort.IntsAreSorted(ints)

	fmt.Println("ints:", ints)
	fmt.Println("sorted?", is)
}
