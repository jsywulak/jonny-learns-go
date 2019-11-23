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
	assert.Equal(t, []string{"a", "b", "c"}, strs)
	ss := sort.StringsAreSorted(strs)

	fmt.Println("Strings:", strs)
	fmt.Println("sorted?", ss)

	ints := []int{7, 2, 4}
	is := sort.IntsAreSorted(ints)

	fmt.Println("ints:", ints)
	fmt.Println("sorted?", is)

	sort.Ints(ints)
	is = sort.IntsAreSorted(ints)

	fmt.Println("ints:", ints)
	fmt.Println("sorted?", is)
}
